package bot_controller

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func (c *BotController) StartHandler(b *gotgbot.Bot, ctx *ext.Context) error {
	err := c.service.Start(b, ctx)
	if err != nil {
		c.logger.Error("Failed to execute Start command: " + err.Error())
		return err
	}
	return nil
}

func (c *BotController) SupportStartHandler(b *gotgbot.Bot, ctx *ext.Context) error {
	err := c.service.SupportStart(b, ctx)
	if err != nil {
		c.logger.Error("Failed to execute SupportStart command: " + err.Error())
		return err
	}
	return nil
}

func (c *BotController) SupportReplyHandler(b *gotgbot.Bot, ctx *ext.Context) error {
	err := c.service.SupportReply(b, ctx)
	if err != nil {
		c.logger.Error("Failed to execute SupportReply command: " + err.Error())
		return err
	}
	return nil
}

func (c *BotController) SendAdminResponseHandler(b *gotgbot.Bot, ctx *ext.Context) error {
	err := c.service.SendAdminResponse(b, ctx)
	if err != nil {
		c.logger.Error("Failed to execute SendAdminResponse command: " + err.Error())
		return err
	}
	return nil
}

func (c *BotController) InlineQueryHandler(b *gotgbot.Bot, ctx *ext.Context) error {
	err := c.service.InlineQuery(b, ctx)
	if err != nil {
		c.logger.Error("Failed to execute InlineQuery command: " + err.Error())
		return err
	}
	return nil
}

func (c *BotController) PreCheckoutHandler(b *gotgbot.Bot, ctx *ext.Context) error {
	err := c.service.PreCheckout(b, ctx)
	if err != nil {
		c.logger.Error("Failed to execute PreCheckout command: " + err.Error())
		return err
	}
	return nil
}

func (c *BotController) PaymentCompleteHandler(b *gotgbot.Bot, ctx *ext.Context) error {
	err := c.service.PaymentComplete(b, ctx)
	if err != nil {
		c.logger.Error("Failed to execute PaymentComplete command: " + err.Error())
		return err
	}
	return nil
}
