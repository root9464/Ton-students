package bot_module

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/precheckoutquery"
	"github.com/root9464/Ton-students/config"
	bot_controller "github.com/root9464/Ton-students/module/bot/controller"
	"github.com/root9464/Ton-students/shared/logger"
)

type BotModule struct {
	bot        *gotgbot.Bot
	dispatcher *ext.Dispatcher
	updater    *ext.Updater
	config     *config.Config
	logger     *logger.Logger
	controller *bot_controller.BotController
}

func NewBotModule(config *config.Config, logger *logger.Logger) (*BotModule, error) {
	bot, err := gotgbot.NewBot(config.BotToken, nil)
	if err != nil {
		logger.Error("failed to create new bot: " + err.Error())
		return nil, err
	}

	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			logger.Error("An error occurred while handling update: " + err.Error())
			return ext.DispatcherActionNoop
		},
		MaxRoutines: ext.DefaultMaxRoutines,
	})

	updater := ext.NewUpdater(dispatcher, nil)
	controller := bot_controller.NewBotController(config, logger)

	return &BotModule{
		bot:        bot,
		dispatcher: dispatcher,
		updater:    updater,
		config:     config,
		logger:     logger,
		controller: controller,
	}, nil
}

func (m *BotModule) Start() error {
	m.registerHandlers()

	if err := m.startPolling(); err != nil {
		return err
	}

	m.logger.Info("👾 Bot started successfully")

	m.updater.Idle()
	return nil
}

func (m *BotModule) registerHandlers() {
	m.dispatcher.AddHandler(handlers.NewCommand("start", m.controller.Start))
	m.dispatcher.AddHandler(handlers.NewCommand("support", m.controller.SupportStart))

	m.dispatcher.AddHandler(handlers.NewCallback(
		filters.CallbackQuery(func(query *gotgbot.CallbackQuery) bool {
			return query.Data != "" && query.Data[:6] == "reply_"
		}),
		m.controller.SupportReply,
	))

	m.dispatcher.AddHandler(handlers.NewMessage(
		filters.Message(func(msg *gotgbot.Message) bool {
			return msg.Chat.Id == m.config.AdminId
		}),
		m.controller.SendAdminResponse,
	))

	m.dispatcher.AddHandler(handlers.NewPreCheckoutQuery(precheckoutquery.All, m.controller.PreCheckout))
	m.dispatcher.AddHandler(handlers.NewMessage(message.SuccessfulPayment, m.controller.PaymentComplete))

	m.dispatcher.AddHandler(handlers.NewInlineQuery(
		filters.InlineQuery(func(query *gotgbot.InlineQuery) bool {
			return query.Query == "invite"
		}),
		m.controller.InlineQuery,
	))
}

func (m *BotModule) startPolling() error {
	err := m.updater.StartPolling(m.bot, &ext.PollingOpts{
		DropPendingUpdates: false,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			Timeout: 60,
		},
	})
	if err != nil {
		m.logger.Error("Failed to start polling: " + err.Error())
		return err
	}
	return nil
}
