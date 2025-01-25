package bot_command

import (
	"sync"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/go-playground/validator/v10"
	"github.com/root9464/Ton-students/config"
	"github.com/root9464/Ton-students/shared/logger"
)

var _ IBotCommand = (*botCommand)(nil)

var replyState = struct {
	mu     sync.Mutex
	active map[int64]int64
}{
	active: make(map[int64]int64),
}

type IBotCommand interface {
	StartMessage(bot *gotgbot.Bot, ctx *ext.Context) error
	GeneratePayment(b *gotgbot.Bot, id string) (*string, error)
	PreCheckout(b *gotgbot.Bot, ctx *ext.Context) error
	PaymentComplete(b *gotgbot.Bot, ctx *ext.Context) error
	InlineQuery(b *gotgbot.Bot, ctx *ext.Context) error
	SupportStart(b *gotgbot.Bot, ctx *ext.Context) error
	SupportReply(b *gotgbot.Bot, ctx *ext.Context) error
	SendAdminResponse(b *gotgbot.Bot, ctx *ext.Context) error
}

type botCommand struct {
	logger    *logger.Logger
	validator *validator.Validate
	config    *config.Config
}

func NewBotCommand(logger *logger.Logger, validator *validator.Validate, config *config.Config) *botCommand {
	return &botCommand{logger: logger, validator: validator, config: config}
}
