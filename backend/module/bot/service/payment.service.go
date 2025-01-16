package bot_service

import (
	"strconv"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func (c *BotService) GeneratePayment(b *gotgbot.Bot, id string) (string, error) {
	//payload := strconv.FormatInt(b.User.Id , 10)
	//log.Println(payload)
	paymentLink, err := b.CreateInvoiceLink("title", "desription", id, "XTR", []gotgbot.LabeledPrice{{
		Label:  "Some product",
		Amount: 1, //1 stars
	}}, nil)

	if err != nil {
		c.logger.Error("Failed to execute GeneratePayment command: " + err.Error())
		return "", err
	}

	return paymentLink, nil
}

func (c *BotService) PreCheckout(b *gotgbot.Bot, ctx *ext.Context) error {

	userId := strconv.FormatInt(ctx.PreCheckoutQuery.From.Id, 10)
	c.logger.Infof("Пользователь, отправивший запрос", userId)

	payload := ctx.PreCheckoutQuery.InvoicePayload
	c.logger.Infof("payload", payload)

	if payload != userId {
		_, err := ctx.PreCheckoutQuery.Answer(b, false, nil)
		if err != nil {
			c.logger.Errorf("failed to answer precheckout query: %w", err)
			return err
		}
		return nil
	}

	_, err := ctx.PreCheckoutQuery.Answer(b, true, nil)
	if err != nil {
		c.logger.Errorf("failed to answer precheckout query: %v", err)
		return err
	}
	return nil
}

func (c *BotService) PaymentComplete(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.EffectiveMessage.Reply(b, "Payment complete - in a real bot, this is where you would provision the product that has been paid for.", nil)
	if err != nil {
		c.logger.Errorf("failed to send payment complete message: %v", err)
		return err
	}
	return nil
}
