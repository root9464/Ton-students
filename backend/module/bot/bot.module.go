package bot_module

import (
	"strings"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
	"github.com/root9464/Ton-students/config"
	bot_controller "github.com/root9464/Ton-students/module/bot/controller"
	"github.com/root9464/Ton-students/shared/logger" // Подключение вашего логгера
)

type BotModule struct {
	bot        *gotgbot.Bot
	dispatcher *ext.Dispatcher
	updater    *ext.Updater
	config     *config.Config
	logger     *logger.Logger // Используем ваш логгер
	controller *bot_controller.BotController
}

func NewBotModule(config *config.Config, logger *logger.Logger) (*BotModule, error) {
	// Создание бота с вашим токеном
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
