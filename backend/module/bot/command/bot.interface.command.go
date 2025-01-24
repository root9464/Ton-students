package bot_command

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/go-playground/validator/v10"
	"github.com/root9464/Ton-students/shared/logger"
)

var _ IBotCommand = (*botCommand)(nil)

type IBotCommand interface {
	StartMessage(bot *gotgbot.Bot, ctx *ext.Context) error
	GeneratePayment(b *gotgbot.Bot, id int64) (*string, error)
	PreCheckout(b *gotgbot.Bot, ctx *ext.Context) error
	PaymentComplete(b *gotgbot.Bot, ctx *ext.Context) error
}

type botCommand struct {
	logger    *logger.Logger
	validator *validator.Validate
}

func NewBotCommand(logger *logger.Logger, validator *validator.Validate) *botCommand {
	return &botCommand{logger: logger, validator: validator}
}
