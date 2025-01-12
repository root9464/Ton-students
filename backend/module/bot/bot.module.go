package bot_module

import (
	"fmt"
	"strings"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/precheckoutquery"
	"github.com/gofiber/fiber/v2"
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

	// Настройка диспетчера с обработкой ошибок
	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			logger.Error("an error occurred while handling update: " + err.Error())
			return ext.DispatcherActionNoop
		},
		MaxRoutines: ext.DefaultMaxRoutines,
	})

	// Создание обновлений
	updater := ext.NewUpdater(dispatcher, nil)

	// Создаем контроллер и передаем ему ваш логгер
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
	//start
	m.dispatcher.AddHandler(handlers.NewCommand("start", m.controller.Start))

	//support
	m.dispatcher.AddHandler(handlers.NewCommand("support", m.controller.SupportStart))

	//reply on button(admin)
	m.dispatcher.AddHandler(handlers.NewCallback(
		filters.CallbackQuery(func(query *gotgbot.CallbackQuery) bool {
			return strings.HasPrefix(query.Data, "reply_")
		}),
		m.controller.SupportReply))

	//response user from admin
	m.dispatcher.AddHandler(handlers.NewMessage(
		filters.Message(func(msg *gotgbot.Message) bool {
			return msg.Chat.Id == m.config.AdminId
		}),
		m.controller.SendAdminResponse))

	//////////////payment
	m.dispatcher.AddHandler(handlers.NewCommand("payment", m.controller.Payment))
	// PreCheckout to handle the step right before payment. Must be handled within 10s, or the checkout will be abandoned by telegram.
	m.dispatcher.AddHandler(handlers.NewPreCheckoutQuery(precheckoutquery.All, preCheckout))
	// Payment received; send/provide product to customer.
	m.dispatcher.AddHandler(handlers.NewMessage(message.SuccessfulPayment, paymentComplete))

	// Запуск обновлений
	err := m.updater.StartPolling(m.bot, &ext.PollingOpts{
		DropPendingUpdates: false,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			Timeout: 60,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 60,
			},
		},
	})
	if err != nil {
		m.logger.Error("failed to start polling: " + err.Error())
		return err
	}

	// Логируем успешный запуск
	m.logger.Info("👾 Bot started successfully")

	m.updater.Idle()
	return nil
}

func (m *BotModule) BotRoutes(router fiber.Router) {
    router.Get("/generate-payment", func(ctx *fiber.Ctx) error {
        return m.controller.GeneratePaymentHandler(m.bot, ctx)
    })
}

func preCheckout(b *gotgbot.Bot, ctx *ext.Context) error {
	// Do any required preCheckout validation here. If anything failed, we should answer the query with "ok: False",
	// and populate the ErrorMessage field in the opts.
	// For example, you may want to ensure that the user who requested the invoice is the same person as the person who
	// is checking out; but this would require storage, so isn't shown here.

	// Answer true once checks have passed.
	_, err := ctx.PreCheckoutQuery.Answer(b, true, nil)
	if err != nil {
		return fmt.Errorf("failed to answer precheckout query: %w", err)
	}
	return nil
}

func paymentComplete(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.EffectiveMessage.Reply(b, "Payment complete - in a real bot, this is where you would provision the product that has been paid for.", nil)
	if err != nil {
		return fmt.Errorf("failed to send payment complete message: %w", err)
	}
	return nil
}
