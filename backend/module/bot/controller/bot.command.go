package bot_controller

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/root9464/Ton-students/config"
	"github.com/root9464/Ton-students/shared/logger"

	service "github.com/root9464/Ton-students/module/bot/service"
)

type BotController struct {
	logger  *logger.Logger
	service *service.BotService
}

func NewBotController(config *config.Config, logger *logger.Logger) *BotController {
	return &BotController{
		service: service.NewBotService(config, logger),
	}
}

//start
func (c *BotController) Start(b *gotgbot.Bot, ctx *ext.Context) error {
	err := c.service.Start(b, ctx)
	if err != nil {
		c.logger.Error("Failed to execute Start command: " + err.Error())
		return err
	}
	return nil
}

///support
func (c *BotController) SupportStart(b *gotgbot.Bot, ctx *ext.Context) error {
	err := c.service.SupportStart(b, ctx)
	if err != nil {
		c.logger.Error("Failed to execute SupportStart command: " + err.Error())
		return err
	}
	return nil
}

func (c *BotController) SupportReply(b *gotgbot.Bot, ctx *ext.Context) error {
	err := c.service.SupportReply(b, ctx)
	if err != nil {
		c.logger.Error("Failed to execute SupportReply command: " + err.Error())
		return err
	}
	return nil
}

func (c *BotController) SendAdminResponse(b *gotgbot.Bot, ctx *ext.Context) error {
	err := c.service.SendAdminResponse(b, ctx)
	if err != nil {
		c.logger.Error("Failed to execute SendAdminResponse command: " + err.Error())
		return err
	}
	return nil
}

//inline 
func (c *BotController) InlineQuery(b *gotgbot.Bot, ctx *ext.Context) error {
	err := c.service.InlineQuery(b, ctx)
	if err != nil {
		c.logger.Error("Failed to execute InlineQuery command: " + err.Error())
		return err
	}
	return nil
}

//payment
func (c *BotController) PreCheckout(b *gotgbot.Bot, ctx *ext.Context) error {
	err := c.service.PreCheckout(b, ctx)
	if err != nil {
		c.logger.Error("Failed to execute PreCheckout command: " + err.Error())
		return err
	}
	return nil
}

func (c *BotController) PaymentComplete(b *gotgbot.Bot, ctx *ext.Context) error {
	err := c.service.PaymentComplete(b, ctx)
	if err != nil {
		c.logger.Error("Failed to execute PaymentComplete command: " + err.Error())
		return err
	}
	return nil
}
