package bot_controller

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/gofiber/fiber/v2"
	"github.com/root9464/Ton-students/shared/logger"

	service "github.com/root9464/Ton-students/module/bot/service"
)

type IBotController interface {
	StartHandler(bot *gotgbot.Bot, ctx *ext.Context) error
	SupportStartHandler(bot *gotgbot.Bot, ctx *ext.Context) error
	SupportReplyHandler(bot *gotgbot.Bot, ctx *ext.Context) error
	SendAdminResponseHandler(bot *gotgbot.Bot, ctx *ext.Context) error
	InlineQueryHandler(bot *gotgbot.Bot, ctx *ext.Context) error
	PreCheckoutHandler(bot *gotgbot.Bot, ctx *ext.Context) error
	PaymentCompleteHandler(bot *gotgbot.Bot, ctx *ext.Context) error

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
