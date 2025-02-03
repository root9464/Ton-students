package bot_controller

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/gofiber/fiber/v2"
	bot_command "github.com/root9464/Ton-students/module/bot/command"
	"github.com/root9464/Ton-students/shared/logger"
)

type IBotController interface {
	Payment(ctx *fiber.Ctx) error
}

type botController struct {
	bot        *gotgbot.Bot
	botCommand bot_command.IBotCommand
	// repo     serv_repository.IServiceModuleRepository
}

func NewBotController(logger *logger.Logger, bot *gotgbot.Bot, botCommand bot_command.IBotCommand) *botController {
	return &botController{
		bot:        bot,
		botCommand: botCommand,
	}
}
