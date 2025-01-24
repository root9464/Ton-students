package bot_command

import (
	"fmt"
	"strconv"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

const (
	PRICE_PAYMENT       = 1
	DESCRIPTION_PAYMENT = `
Превратите знания в доход! 🎓💰
Покупайте и продавайте курсовые, дипломные и другие студенческие работы. 
Найдите готовые решения или предлагайте свои услуги другим студентам. 
Подключите премиум-подписку, чтобы зарабатывать на своих знаниях без ограничений! 🚀
	`
)

func (cm *botCommand) StartMessage(bot *gotgbot.Bot, ctx *ext.Context) error {
	username := ctx.EffectiveUser.FirstName
	if username == "" {
		username = "there"
	}

	messageText := fmt.Sprintf("Hello, %s! 👋\nWelcome to the bot. How can I assist you today?", username)

	_, err := ctx.EffectiveMessage.Reply(bot, messageText, nil)
	if err != nil {
		cm.logger.Error("Failed to send start message: " + err.Error())
		return err
	}

	return nil
}

func (cm *botCommand) GeneratePayment(b *gotgbot.Bot, id int64) (*string, error) {
	cm.logger.Infof("Generating payment for id: %v", id)

	if id == 0 {
		return nil, fmt.Errorf("id cannot be 0")
	}

	cm.logger.Infof("Validation successful: %v", id)

	cm.logger.Info("Creating invoice link...")
	paymentLink, err := b.CreateInvoiceLink("Creator subscription", DESCRIPTION_PAYMENT, fmt.Sprint(id), "XTR", []gotgbot.LabeledPrice{{
		Label:  "TonStudents",
		Amount: PRICE_PAYMENT,
	}}, &gotgbot.CreateInvoiceLinkOpts{
		PhotoUrl: "https://cdn-icons-png.flaticon.com/512/4689/4689222.png",
	})

	if err != nil {
		cm.logger.Error("Failed to execute GeneratePayment command: " + err.Error())
		return nil, err
	}

	cm.logger.Infof("Invoice link created successfully: %s", paymentLink)
	return &paymentLink, nil
}

func (cm *botCommand) PreCheckout(b *gotgbot.Bot, ctx *ext.Context) error {

	userId := strconv.FormatInt(ctx.PreCheckoutQuery.From.Id, 10)
	cm.logger.Infof("user id %s", userId)

	payload := ctx.PreCheckoutQuery.InvoicePayload
	cm.logger.Infof("getting payload %v", payload)

	if payload != userId {
		_, err := ctx.PreCheckoutQuery.Answer(b, false, nil)
		if err != nil {
			cm.logger.Errorf("failed to answer precheckout query: %s", err)
			return err
		}
		return nil
	}

	_, err := ctx.PreCheckoutQuery.Answer(b, true, nil)
	if err != nil {
		cm.logger.Errorf("failed to answer precheckout query: %v", err)
		return err
	}
	return nil
}

func (cm *botCommand) PaymentComplete(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.EffectiveMessage.Reply(b, "Payment complete - in a real bot, this is where you would provision the product that has been paid for.", nil)
	if err != nil {
		cm.logger.Errorf("failed to send payment complete message: %v", err)
		return err
	}
	return nil
}
