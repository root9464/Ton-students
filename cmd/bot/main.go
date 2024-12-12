package main

import (
	"log"
	"root/pkg/config"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"

	command "root/internal/bot/handler"
)

func main() {
	cfg := config.New()

	b, err := gotgbot.NewBot(cfg.BotToken, nil)
	if err != nil {
		panic("failed to create new bot: " + err.Error())
	}

	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		// If an error is returned by a handler, log it and continue going.
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			log.Println("an error occurred while handling update:", err.Error())
			return ext.DispatcherActionNoop
		},
		MaxRoutines: ext.DefaultMaxRoutines,
	})
	updater := ext.NewUpdater(dispatcher, nil)


	dispatcher.AddHandler(handlers.NewCommand("start", command.Start))

	err = updater.StartPolling(b, &ext.PollingOpts{
		DropPendingUpdates: true,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			Timeout: 10,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 10,
			},
		},
	})

	if err != nil {
		panic("failed to start polling: " + err.Error())
	}

	log.Println("bot started")

	updater.Idle()
}
