package bot_service

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/root9464/Ton-students/config"
	"github.com/root9464/Ton-students/shared/logger"
)

type IBotService interface {
	Start(b *gotgbot.Bot, ctx *ext.Context) error

	SupportReply(b *gotgbot.Bot, ctx *ext.Context) error
	SupportStart(b *gotgbot.Bot, ctx *ext.Context) error
	SendAdminResponse(b *gotgbot.Bot, ctx *ext.Context) error

	GeneratePayment(b *gotgbot.Bot, id string) (string, error) 
	PreCheckout(b *gotgbot.Bot, ctx *ext.Context) error
	PaymentComplete(b *gotgbot.Bot, ctx *ext.Context) error

	InvateLink(b *gotgbot.Bot) (string, error)
	InlineQuery(b *gotgbot.Bot, ctx *ext.Context) error
}

type BotService struct {
	config *config.Config
	logger *logger.Logger
}

func NewBotService(config *config.Config, logger *logger.Logger) *BotService {
	return &BotService{config: config, logger: logger}
}
