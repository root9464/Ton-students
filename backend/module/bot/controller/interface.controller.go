package bot_controller

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/gofiber/fiber/v2"
	"github.com/root9464/Ton-students/shared/logger"

	service "github.com/root9464/Ton-students/module/bot/service"
)

type IBotController interface {
	Start(bot *gotgbot.Bot, ctx *ext.Context) error
	SupportStart(bot *gotgbot.Bot, ctx *ext.Context) error
	SupportReply(bot *gotgbot.Bot, ctx *ext.Context) error
	SendAdminResponse(bot *gotgbot.Bot, ctx *ext.Context) error
	InlineQuery(bot *gotgbot.Bot, ctx *ext.Context) error
	PreCheckout(bot *gotgbot.Bot, ctx *ext.Context) error
	PaymentComplete(bot *gotgbot.Bot, ctx *ext.Context) error

	GeneratePaymentHandler(ctx *fiber.Ctx) error
}

type BotController struct {
	bot     *gotgbot.Bot
	service service.IBotService
	logger  *logger.Logger
}

func NewBotController(bot *gotgbot.Bot, service service.IBotService, logger *logger.Logger) *BotController {
	return &BotController{bot: bot, service: service, logger: logger}
}
