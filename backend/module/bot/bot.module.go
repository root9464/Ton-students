package bot_module

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/precheckoutquery"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/root9464/Ton-students/config"
	bot_command "github.com/root9464/Ton-students/module/bot/command"
	bot_controller "github.com/root9464/Ton-students/module/bot/controller"
	jwt_module "github.com/root9464/Ton-students/module/jwt"
	user_module "github.com/root9464/Ton-students/module/user"
	"github.com/root9464/Ton-students/shared/logger"
	"github.com/root9464/Ton-students/shared/middleware"
	"gorm.io/gorm"
)

type BotModule struct {
	bot        *gotgbot.Bot
	dispatcher *ext.Dispatcher
	updater    *ext.Updater

	botCommand    bot_command.IBotCommand
	botController bot_controller.IBotController

	logger    *logger.Logger
	validator *validator.Validate
	db        *gorm.DB
	config    *config.Config

	userModule user_module.UserModule
	jwtModule  jwt_module.JwtModule
}

func NewBotModule(
	logger *logger.Logger,
	validator *validator.Validate,
	db *gorm.DB,
	config *config.Config,

	userModule user_module.UserModule,
	jwtModule jwt_module.JwtModule,
) *BotModule {
	bot, err := gotgbot.NewBot(config.TelegramBotToken, &gotgbot.BotOpts{
		BotClient: &gotgbot.BaseBotClient{
			UseTestEnvironment: true,
		},
	})
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

func (m *BotModule) registerCommands() {
	m.dispatcher.AddHandler(handlers.NewCommand("start", m.BotCommand().StartMessage))

	m.dispatcher.AddHandler(handlers.NewPreCheckoutQuery(precheckoutquery.All, m.BotCommand().PreCheckout))
	m.dispatcher.AddHandler(handlers.NewMessage(message.SuccessfulPayment, m.BotCommand().PaymentComplete))

	m.dispatcher.AddHandler(handlers.NewCommand("support", m.BotCommand().SupportStart))
	m.dispatcher.AddHandler(handlers.NewCallback(
		filters.CallbackQuery(func(query *gotgbot.CallbackQuery) bool {
			return query.Data != "" && query.Data[:6] == "reply_"
		}),
		m.BotCommand().SupportReply,
	))
	m.dispatcher.AddHandler(handlers.NewMessage(
		filters.Message(func(msg *gotgbot.Message) bool {
			return msg.Chat.Id == m.config.AdminId
		}),
		m.BotCommand().SendAdminResponse,
	))

	m.dispatcher.AddHandler(handlers.NewInlineQuery(
		filters.InlineQuery(func(query *gotgbot.InlineQuery) bool {
			return query.Query == "invite"
		}),
		m.BotCommand().InlineQuery,
	))

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

func (m *BotModule) BotCommand() bot_command.IBotCommand {
	if m.botCommand == nil {
		m.botCommand = bot_command.NewBotCommand(m.logger, m.validator, m.config)
	}
	return m.botCommand
}

func (m *BotModule) BotController() bot_controller.IBotController {
	if m.botController == nil {
		m.botController = bot_controller.NewBotController(m.logger, m.bot, m.BotCommand())
	}
	return m.botController
}

func (m *BotModule) BotRoutes(router fiber.Router) {

	middleware := middleware.NewMiddleware(m.logger, m.userModule.UserRepo(), m.jwtModule.JwtHelpers(), m.config.JwtPublicKey)

	bot := router.Group("/bot", middleware.UserOnly())

	bot.Get("/payment", m.BotController().Payment)
}
