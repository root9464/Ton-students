package bot_controller

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/root9464/Ton-students/shared/logger" // Подключаем ваш логгер

	"github.com/root9464/Ton-students/config"
)

type BotController struct {
	config *config.Config
	logger *logger.Logger // Используем ваш логгер
}

func NewBotController(config *config.Config, logger *logger.Logger) *BotController {
	return &BotController{
		config: config,
		logger: logger,
	}
}

// Start - обработчик команды /start
func (c *BotController) Start(b *gotgbot.Bot, ctx *ext.Context) error {
	c.logger.Info("Start command received")
	_, err := b.SendMessage(ctx.EffectiveChat.Id, "Hello! How can I help you?", nil)
	if err != nil {
		c.logger.Error("Error sending start message: " + err.Error())
	}
	return err
}

// SupportStart - обработчик команды /support
func (c *BotController) SupportStart(b *gotgbot.Bot, ctx *ext.Context) error {
	c.logger.Info("SupportStart command received")
	_, err := b.SendMessage(ctx.EffectiveChat.Id, "Please describe your issue.", nil)
	if err != nil {
		c.logger.Error("Error sending support message: " + err.Error())
	}
	return err
}

// SupportReply - обработчик для ответа на запросы поддержки
func (c *BotController) SupportReply(b *gotgbot.Bot, ctx *ext.Context) error {
	c.logger.Info("SupportReply called")
	_, err := b.SendMessage(ctx.EffectiveChat.Id, "Support response: Here is your answer.", nil)
	if err != nil {
		c.logger.Error("Error sending support reply: " + err.Error())
	}
	return err
}
