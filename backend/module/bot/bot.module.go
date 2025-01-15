package bot_module

import (
	"fmt"
	"log"
	"strconv"
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
	//m.dispatcher.AddHandler(handlers.NewCommand("payment", m.controller.Payment))
	// PreCheckout to handle the step right before payment. Must be handled within 10s, or the checkout will be abandoned by telegram.
	m.dispatcher.AddHandler(handlers.NewPreCheckoutQuery(precheckoutquery.All, preCheckout))
	// Payment received; send/provide product to customer.
	m.dispatcher.AddHandler(handlers.NewMessage(message.SuccessfulPayment, paymentComplete))

	//  inline-запросы
	m.dispatcher.AddHandler(handlers.NewInlineQuery(
		filters.InlineQuery(func(query *gotgbot.InlineQuery) bool {
			// Опционально фильтруем запросы, например, по наличию текста "invite"
			return strings.Contains(query.Query, "invite")
		}),
		handleInlineQuery, // Обработчик для inline-запросов
	))
	
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

	router.Get("/ref", func(ctx *fiber.Ctx) error {
		return m.controller.InvateLinkHandler(m.bot, ctx)
	})
}

func preCheckout(b *gotgbot.Bot, ctx *ext.Context) error {

	userId := strconv.FormatInt(ctx.PreCheckoutQuery.From.Id, 10)
	log.Println("Пользователь, отправивший запрос", userId)

	payload := ctx.PreCheckoutQuery.InvoicePayload
	log.Println("payload", payload)

	if payload != userId {
		// Answer false to cancel the checkout.
		_, err := ctx.PreCheckoutQuery.Answer(b, false, nil)
		if err != nil {
			return fmt.Errorf("failed to answer precheckout query: %w", err)
		}
		return nil
	}

	// Answer true once checks have passed.
	_, err := ctx.PreCheckoutQuery.Answer(b, true, nil)
	if err != nil {
		return fmt.Errorf("failed to answer precheckout query: %w", err)
	}
	//log.Println(ctx.PreCheckoutQuery.Id)
	return nil
}

func paymentComplete(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.EffectiveMessage.Reply(b, "Payment complete - in a real bot, this is where you would provision the product that has been paid for.", nil)
	if err != nil {
		return fmt.Errorf("failed to send payment complete message: %w", err)
	}
	return nil
}

func handleInlineQuery(b *gotgbot.Bot, ctx *ext.Context) error {
	// Получаем текст запроса пользователя
	query := ctx.InlineQuery.Query

	// Проверяем, что запрос содержит "invite"
	if query == "invite" {
		// Создаем inline-результат с кнопкой
		results := []gotgbot.InlineQueryResult{
			&gotgbot.InlineQueryResultArticle{
				Id:    "1", // уникальный ID результата
				Title: "Send Invite Link",
				InputMessageContent: &gotgbot.InputTextMessageContent{
					MessageText: "Click the button below to join!",
				},
				ReplyMarkup: &gotgbot.InlineKeyboardMarkup{
					InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
						{
							{
								Text: "Join Now", // Текст кнопки
								Url:  "https://t.me/ttonstudents_bot?start=12345678", // Ссылка на приглашение
							},
						},
					},
				},
			},
		}

		// Отправляем inline-результаты
		_, err := b.AnswerInlineQuery(ctx.InlineQuery.Id, results, &gotgbot.AnswerInlineQueryOpts{
			CacheTime: 0, // Обнуляем кэш, чтобы изменения были видны сразу
		})
		if err != nil {
			return fmt.Errorf("failed to answer inline query: %w", err)
		}
	}

	return nil
}
