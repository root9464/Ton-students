package bot_controller

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/root9464/Ton-students/shared/logger"

	service "github.com/root9464/Ton-students/module/bot/service"
)

type IBotController interface {
	Start (ctx *fiber.Ctx) error


	InvateLinkHandler(bot *gotgbot.Bot, ctx *fiber.Ctx) error
	GeneratePaymentHandler(bot *gotgbot.Bot, ctx *fiber.Ctx) error
}

type BotController struct {
	bot     *gotgbot.Bot
	service service.IBotService
	logger  *logger.Logger
}

func NewBotController(bot *gotgbot.Bot, service service.IBotService, logger *logger.Logger) *BotController {
	return &BotController{bot: bot, service: service, logger: logger}
}
