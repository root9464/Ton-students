package bot_module

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/precheckoutquery"
	"github.com/gofiber/fiber/v2"
	"github.com/root9464/Ton-students/config"
	controller "github.com/root9464/Ton-students/module/bot/controller"
	service "github.com/root9464/Ton-students/module/bot/service"
	"github.com/root9464/Ton-students/shared/logger"
)

type BotModule struct {
	bot           *gotgbot.Bot
	dispatcher    *ext.Dispatcher
	updater       *ext.Updater
	config        *config.Config
	logger        *logger.Logger
	botController controller.IBotController
	botService    service.IBotService
}

func NewBotModule(config *config.Config, logger *logger.Logger) (*BotModule, error) {
	bot, err := gotgbot.NewBot(config.TelegramBotToken, nil)
	if err != nil {
		logger.Error("failed to create new bot: " + err.Error())
		return nil, err
	}

	return &BotModule{
		bot:    bot,
		config: config,
		logger: logger,
	}, nil
}

func (m *BotModule) BotService() service.IBotService {
	if m.botService == nil {
		m.botService = service.NewBotService(m.config, m.logger)
	}
	return m.botService
}

func (m *BotModule) BotController() controller.IBotController {
	if m.botController == nil {
		m.botController = controller.NewBotController(m.bot, m.BotService(), m.logger)
	}
	return m.botController
}

func (m *BotModule) InitBot() error {
	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			m.logger.Error("An error occurred while handling update: " + err.Error())
			return ext.DispatcherActionNoop
		},
		MaxRoutines: ext.DefaultMaxRoutines,
	})

	updater := ext.NewUpdater(dispatcher, nil)
	err := updater.StartPolling(m.bot, &ext.PollingOpts{
		DropPendingUpdates: true,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			Timeout: 60,
		},
	})

	if err != nil {
		m.logger.Error("Failed to start polling: " + err.Error())
		return err
	}

	m.dispatcher = dispatcher
	m.updater = updater

	m.registerHandlers()

	m.logger.Info("👾 Bot started successfully")
	m.updater.Idle()
	return nil
}

func (m *BotModule) registerHandlers() {
	m.dispatcher.AddHandler(handlers.NewCommand("start", m.BotController().StartHandler))
	m.dispatcher.AddHandler(handlers.NewCommand("support", m.BotController().SupportStartHandler))

	m.dispatcher.AddHandler(handlers.NewCallback(
		filters.CallbackQuery(func(query *gotgbot.CallbackQuery) bool {
			return query.Data != "" && query.Data[:6] == "reply_"
		}),
		m.BotController().SupportReplyHandler,
	))

	m.dispatcher.AddHandler(handlers.NewMessage(
		filters.Message(func(msg *gotgbot.Message) bool {
			return msg.Chat.Id == m.config.AdminId
		}),
		m.BotController().SendAdminResponseHandler,
	))

	m.dispatcher.AddHandler(handlers.NewPreCheckoutQuery(precheckoutquery.All, m.BotController().PreCheckoutHandler))
	m.dispatcher.AddHandler(handlers.NewMessage(message.SuccessfulPayment, m.BotController().PaymentCompleteHandler))

	m.dispatcher.AddHandler(handlers.NewInlineQuery(
		filters.InlineQuery(func(query *gotgbot.InlineQuery) bool {
			return query.Query == "invite"
		}),
		m.BotController().InlineQueryHandler,
	))
}

func (m *BotModule) BotRoutes(router fiber.Router) {
	router.Get("/generate-payment", m.BotController().GeneratePaymentHandler)
}
