package bot_controller

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/root9464/Ton-students/shared/logger" // Подключаем ваш логгер

	"github.com/root9464/Ton-students/config"
	service "github.com/root9464/Ton-students/module/bot/service"
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

func (c *BotController) Start(b *gotgbot.Bot, ctx *ext.Context) error {
	err := service.Start(b, ctx, c.config.ChannelId, c.logger.Logger)
	if err != nil {
		c.logger.Error("Failed to execute Start command: " + err.Error())
		return err
	}
	return nil
}

func (c *BotController) SupportStart(b *gotgbot.Bot, ctx *ext.Context) error {
	err := service.SupportStart(b, ctx, c.config.AdminId, c.logger.Logger)
	if err != nil {
		c.logger.Error("Failed to execute SupportStart command: " + err.Error())
		return err
	}
	return nil
}

func (c *BotController) SupportReply(b *gotgbot.Bot, ctx *ext.Context) error {
	err := service.SupportReply(b, ctx, c.logger.Logger)
	if err != nil {
		c.logger.Error("Failed to execute SupportReply command: " + err.Error())
		return err
	}
	return nil
}

func (c *BotController) SendAdminResponse(b *gotgbot.Bot, ctx *ext.Context) error {
	err := service.SendAdminResponse(b, ctx, c.logger.Logger)
	if err != nil {
		c.logger.Error("Failed to execute SendAdminResponse command: " + err.Error())
		return err
	}
	return nil
}
