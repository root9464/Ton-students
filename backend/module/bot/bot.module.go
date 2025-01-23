package bot_module

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/go-playground/validator/v10"
	"github.com/root9464/Ton-students/config"
	bot_command "github.com/root9464/Ton-students/module/bot/command"
	jwt_module "github.com/root9464/Ton-students/module/jwt"
	"github.com/root9464/Ton-students/shared/logger"
	"gorm.io/gorm"
)

type BotModule struct {
	bot        *gotgbot.Bot
	dispatcher *ext.Dispatcher
	updater    *ext.Updater

	botCommand bot_command.IBotCommand

	logger    *logger.Logger
	validator *validator.Validate
	db        *gorm.DB
	config    *config.Config

	jwtModule jwt_module.JwtModule
}

func (m *BotModule) BotCommand() bot_command.IBotCommand {
	if m.botCommand == nil {
		m.botCommand = bot_command.NewBotCommand(m.logger)
	}
	return m.botCommand
}

func (m *BotModule) registerCommands() {
	m.dispatcher.AddHandler(handlers.NewCommand("start", m.BotCommand().StartMessage))
}

func (m *BotModule) InitBot() error {
	m.dispatcher = ext.NewDispatcher(&ext.DispatcherOpts{
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			m.logger.Error("An error occurred while handling update: " + err.Error())
			return ext.DispatcherActionNoop
		},
		MaxRoutines: ext.DefaultMaxRoutines,
	})

	m.updater = ext.NewUpdater(m.dispatcher, nil)
	err := m.updater.StartPolling(m.bot, &ext.PollingOpts{
		DropPendingUpdates: true,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			Limit: 4,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 60,
			},
		},
	})

	if err != nil {
		m.logger.Error("Failed to start polling: " + err.Error())
		return err
	}

	m.registerCommands()

	m.logger.Info("👾 Bot started successfully")
	m.updater.Idle()
	return nil
}

func NewBotModule(logger *logger.Logger, validator *validator.Validate, db *gorm.DB, config *config.Config, jwtModule jwt_module.JwtModule) *BotModule {
	bot, err := gotgbot.NewBot(config.TelegramBotToken, nil)
	if err != nil {
		logger.Error("failed to create new bot: " + err.Error())
		return nil
	}

	return &BotModule{
		bot: bot,

		logger:    logger,
		validator: validator,
		db:        db,
		config:    config,
		jwtModule: jwtModule,
	}
}
