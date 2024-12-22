package bot_model

import (
	"root/internal/bot/service"
	"root/pkg/config"
	"root/pkg/logger"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
)

type BotModule struct {
	bot        *gotgbot.Bot
	updater    *ext.Updater
	dispatcher *ext.Dispatcher
	service    *service.BotService
	log        *logger.Logger
	config     *config.Config
}

func NewBotModule(log *logger.Logger, config *config.Config) *BotModule {
	return &BotModule{
		log:    log,
		config: config,
	}
}

func (p *BotModule) BotModule() error {
	// Initialize Bot
	bot, err := gotgbot.NewBot(p.config.BotToken, nil)
	if err != nil {
		p.log.Panic("failed to create new bot:", err)
	}

	// Initialize the bot service
	p.service = service.NewBotService(bot, p.config, p.log)

	// Dispatcher setup
	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			p.log.Error("Error while handling update:", err)
			return ext.DispatcherActionNoop
		},
		MaxRoutines: ext.DefaultMaxRoutines,
	})

	updater := ext.NewUpdater(dispatcher, nil)

	// Register Handlers
	p.registerHandlers(dispatcher, bot)

	// Start polling for updates
	err = updater.StartPolling(bot, &ext.PollingOpts{
		DropPendingUpdates: false,
	})
	if err != nil {
		p.log.Panic("failed to start polling:", err)
	}

	p.updater = updater
	p.dispatcher = dispatcher
	p.bot = bot

	p.log.Info("Bot started")
	updater.Idle()

	return nil
}

func (p *BotModule) registerHandlers(dispatcher *ext.Dispatcher, bot *gotgbot.Bot) {
	dispatcher.AddHandler(handlers.NewCommand("start", p.service.Start))
	dispatcher.AddHandler(handlers.NewCommand("support", p.service.SupportStart))
	dispatcher.AddHandler(handlers.NewCallback(filters.CallbackQuery(func(query *gotgbot.CallbackQuery) bool {
		return query.Data == "reply"
	}), p.service.SupportReply))
	dispatcher.AddHandler(handlers.NewMessage(filters.Message(func(msg *gotgbot.Message) bool {
		return msg.Chat.Id == p.config.AdminId
	}), p.service.SendAdminResponse))
}
