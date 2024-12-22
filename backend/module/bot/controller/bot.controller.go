package handler

import (
	"root/internal/bot/service"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

type BotHandler struct {
	service *service.BotService
}

func NewBotHandler(service *service.BotService) *BotHandler {
	return &BotHandler{
		service: service,
	}
}

func (h *BotHandler) Start(b *gotgbot.Bot, ctx *ext.Context) error {
	return h.service.Start(b, ctx)
}

func (h *BotHandler) SupportStart(b *gotgbot.Bot, ctx *ext.Context) error {
	return h.service.SupportStart(b, ctx)
}

func (h *BotHandler) SupportReply(b *gotgbot.Bot, ctx *ext.Context) error {
	return h.service.SupportReply(b, ctx)
}

func (h *BotHandler) SendAdminResponse(b *gotgbot.Bot, ctx *ext.Context) error {
	return h.service.SendAdminResponse(b, ctx)
}
