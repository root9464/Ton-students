package handler

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

type SStart struct{}

func Start(b *gotgbot.Bot, ctx *ext.Context) error {
	b.SendMessage(ctx.EffectiveChat.Id, "Hello! How can I help you?", nil)

	return nil
}
