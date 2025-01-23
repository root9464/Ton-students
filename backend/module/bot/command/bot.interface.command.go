package bot_command

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/root9464/Ton-students/shared/logger"
)

type IBotCommand interface {
	StartMessage(bot *gotgbot.Bot, ctx *ext.Context) error
}

type botCommand struct {
	logger *logger.Logger
}

func NewBotCommand(logger *logger.Logger) *botCommand {
	return &botCommand{logger: logger}
}
