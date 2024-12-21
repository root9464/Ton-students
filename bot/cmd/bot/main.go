package main

import (
	"strings"
	"time"

	"root/pkg/config"
	"root/pkg/logger" // Подключение пакета логгера

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"

	command "root/internal/bot/handler"
)

func main() {
	cfg := config.New()
	log := logger.New() 
	
	b, err := gotgbot.NewBot(cfg.BotToken, nil)
	if err != nil {
		log.Panic("failed to create new bot:", err)
	}

	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			log.Error("an error occurred while handling update:", err)
			return ext.DispatcherActionNoop
		},
		MaxRoutines: ext.DefaultMaxRoutines,
	})
	updater := ext.NewUpdater(dispatcher, nil)

	dispatcher.AddHandler(handlers.NewCommand("start", func(b *gotgbot.Bot, ctx *ext.Context) error {
		return command.Start(b, ctx, cfg.ChannelId, log)
	}))

	dispatcher.AddHandler(handlers.NewCommand("support", func(b *gotgbot.Bot, ctx *ext.Context) error {
		return command.SupportStart(b, ctx, cfg.AdminId, log)
	}))

	dispatcher.AddHandler(handlers.NewCallback(filters.CallbackQuery(func(query *gotgbot.CallbackQuery) bool {
		return strings.HasPrefix(query.Data, "reply_")
	}), func(b *gotgbot.Bot, ctx *ext.Context) error {
		return command.SupportReply(b, ctx, log)
	}))

	dispatcher.AddHandler(handlers.NewMessage(filters.Message(func(msg *gotgbot.Message) bool {
		return msg.Chat.Id == cfg.AdminId
	}), func(b *gotgbot.Bot, ctx *ext.Context) error {
		return command.SendAdminResponse(b, ctx, log)
	}))

	err = updater.StartPolling(b, &ext.PollingOpts{
		DropPendingUpdates: false,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			Timeout: 60,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 60,
			},
		},
	})

	if err != nil {
		log.Panic("failed to start polling:", err)
	}

	log.Info("bot started")
	updater.Idle()
}
