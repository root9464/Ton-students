package main

import (
	"log"
	"strings"
	"time"

	"root/pkg/config"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
	"github.com/sirupsen/logrus"

	command "root/internal/bot/handler"
)

func main() {
	cfg := config.New()

	b, err := gotgbot.NewBot(cfg.BotToken, nil)
	if err != nil {
		panic("failed to create new bot: " + err.Error())
	}

	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			log.Println("an error occurred while handling update:", err.Error())
			return ext.DispatcherActionNoop
		},
		MaxRoutines: ext.DefaultMaxRoutines,
	})
	updater := ext.NewUpdater(dispatcher, nil)

	// Обработчик команды /start
	dispatcher.AddHandler(handlers.NewCommand("start", func(b *gotgbot.Bot, ctx *ext.Context) error {
		return trackAndHandle(b, ctx, command.Start)
	}))

	// Обработчик команды /support
	dispatcher.AddHandler(handlers.NewCommand("support", func(b *gotgbot.Bot, ctx *ext.Context) error {
		return command.SupportStart(b, ctx, cfg.AdminId, logrus.New())
	}))

	// Обработчик нажатия кнопки "Ответить"
	dispatcher.AddHandler(handlers.NewCallback(filters.CallbackQuery(func(query *gotgbot.CallbackQuery) bool {
		return strings.HasPrefix(query.Data, "reply_")
	}), func(b *gotgbot.Bot, ctx *ext.Context) error {
		return command.SupportReply(b, ctx, logrus.New())
	}))

	// Обработчик текстовых сообщений администратора
	dispatcher.AddHandler(handlers.NewMessage(filters.Message(func(msg *gotgbot.Message) bool {
		return msg.Chat.Id == cfg.AdminId // Сообщения обрабатываются только от администратора
	}), func(b *gotgbot.Bot, ctx *ext.Context) error {
		return command.SendAdminResponse(b, ctx, logrus.New())
	}))

	// Запуск polling
	err = updater.StartPolling(b, &ext.PollingOpts{
		DropPendingUpdates: true,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			Timeout: 60,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 60,
			},
		},
	})

	if err != nil {
		panic("failed to start polling: " + err.Error())
	}

	log.Println("bot started")

	updater.Idle()
}

func trackAndHandle(b *gotgbot.Bot, ctx *ext.Context, handler func(b *gotgbot.Bot, ctx *ext.Context) error) error {
	// Обрабатываем текущее сообщение
	err := handler(b, ctx)

	return err
}
