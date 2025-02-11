package bot_command

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	keyboards "github.com/root9464/Ton-students/module/bot/keyboards"
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

	_, err := ctx.EffectiveMessage.Reply(bot, messageText, &gotgbot.SendMessageOpts{
		ReplyMarkup: keyboards.Keyboard,
	})
	if err != nil {
		cm.logger.Error("Failed to send start message: " + err.Error())
		return err
	}

	return nil
}

func (cm *botCommand) GeneratePayment(b *gotgbot.Bot, id string) (*string, error) {
	cm.logger.Infof("Generating payment for id: %v", id)

	if id == "" {
		return nil, fmt.Errorf("id cannot be empty")
	}

	cm.logger.Infof("Validation successful: %v", id)

	cm.logger.Info("Creating invoice link...")
	paymentLink, err := b.CreateInvoiceLink("Creator subscription", DESCRIPTION_PAYMENT, id, "XTR", []gotgbot.LabeledPrice{{
		Label:  "TonStudents",
		Amount: PRICE_PAYMENT,
	}}, &gotgbot.CreateInvoiceLinkOpts{
		PhotoUrl:           "https://cdn-icons-png.flaticon.com/512/4689/4689222.png",
		SubscriptionPeriod: 2592000,
	})

	cm.logger.Info("Executing GeneratePayment command...")
	if err != nil {
		cm.logger.Error("Failed to execute GeneratePayment command: " + err.Error())
		return nil, err
	}

	cm.logger.Infof("Invoice link created successfully: %s", paymentLink)
	return &paymentLink, nil
}

func (cm *botCommand) PreCheckout(b *gotgbot.Bot, ctx *ext.Context) error {
	cm.logger.Infof("PreCheckout called with user id: %v", ctx.PreCheckoutQuery.From.Id)

	userId := strconv.FormatInt(ctx.PreCheckoutQuery.From.Id, 10)
	cm.logger.Infof("user id %s", userId)

	payload := ctx.PreCheckoutQuery.InvoicePayload
	cm.logger.Infof("getting payload %v", payload)

	if payload != userId {
		_, err := ctx.PreCheckoutQuery.Answer(b, false, &gotgbot.AnswerPreCheckoutQueryOpts{
			ErrorMessage: "Похоже вы пытаетесь оплатить за другого пользователя.",
		})
		if err != nil {
			cm.logger.Errorf("failed to answer precheckout query: %s", err)
			return err
		}
		return nil
	}

	cm.logger.Infof("Payload matches user id, proceeding with payment...")
	_, err := ctx.PreCheckoutQuery.Answer(b, true, nil)
	if err != nil {
		cm.logger.Errorf("failed to answer precheckout query: %v", err)
		return err
	}

	cm.logger.Info("Payment successful")
	return nil
}

func (cm *botCommand) PaymentComplete(b *gotgbot.Bot, ctx *ext.Context) error {
	cm.logger.Info("Payment complete")

	cm.logger.Infof("Payload: %v", ctx.Message.SuccessfulPayment.InvoicePayload)
	cm.logger.Infof("Date: %v", ctx.Message.SuccessfulPayment.SubscriptionExpirationDate)
	cm.logger.Infof("All: %v", ctx.Message.SuccessfulPayment)

	err := cm.userRepo.ChangeUserRole(context.Background(), ctx.Message.SuccessfulPayment.InvoicePayload, "creator")
	if err != nil {
		cm.logger.Errorf("failed to change user role: %v", err)
		return err
	}

	cm.logger.Infof("Granting privileges to the user... %v", ctx.Message.SuccessfulPayment.OrderInfo)

	_, err = ctx.EffectiveMessage.Reply(b, "Оплата прошла успешно, поздравляем с преобретением подписки.", nil)
	if err != nil {
		cm.logger.Errorf("failed to send payment complete message: %v", err)
		return err
	}
	cm.logger.Info("Payment complete message sent")
	return nil
}

func (cm *botCommand) SupportStart(b *gotgbot.Bot, ctx *ext.Context) error {
	cm.logger.Infof("SupportStart called with userID: %v", ctx.EffectiveUser.Id)

	userID := ctx.EffectiveUser.Id
	args := ctx.Args()

	if len(args) == 0 {
		_, err := ctx.EffectiveMessage.Reply(b,
			"❓ <b>Введите ваш вопрос</b>\n\nПример:\n<code>/support Как зарегистрироваться?</code>",
			&gotgbot.SendMessageOpts{ParseMode: "HTML"})
		return err
	}

	question := strings.Join(args, " ")
	cm.logger.Infof("Received support question from userID %d: %s", userID, question)

	_, err := b.SendMessage(cm.config.AdminId, fmt.Sprintf(
		"📩 <b>Новый запрос от пользователя</b>\n\n<b>Пользователь:</b> @%s\n<b>ID:</b> <code>%d</code>\n\n<b>Вопрос:</b>\n%s",
		ctx.EffectiveUser.Username, userID, question,
	), &gotgbot.SendMessageOpts{
		ParseMode: "HTML",
	})
	if err != nil {
		cm.logger.Errorf("Error sending support message to admin: %v", err)
		return err
	}

	_, err = b.SendMessage(userID,
		"✅ <b>Ваш запрос отправлен в поддержку</b>\n\nПожалуйста, ожидайте ответа от администратора.",
		&gotgbot.SendMessageOpts{ParseMode: "HTML"})
	if err != nil {
		cm.logger.Errorf("Error sending support message to user: %v", err)
		return err
	}

	replyMarkup := &gotgbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
			{
				{
					Text:         "Ответить",
					CallbackData: fmt.Sprintf("reply_%d", userID),
				},
			},
		},
	}
	_, err = b.SendMessage(cm.config.AdminId,
		"👤 <b>Вы можете ответить на запрос этого пользователя:</b>",
		&gotgbot.SendMessageOpts{
			ParseMode:   "HTML",
			ReplyMarkup: replyMarkup,
		})
	return err
}

func (cm *botCommand) InlineQuery(b *gotgbot.Bot, ctx *ext.Context) error {
	query := ctx.InlineQuery.Query

	if query == "invite" {
		results := []gotgbot.InlineQueryResult{
			&gotgbot.InlineQueryResultArticle{
				Id:    "1",
				Title: "Send Invite Link",
				InputMessageContent: &gotgbot.InputTextMessageContent{
					MessageText: "Click the button below to join!",
				},
				ReplyMarkup: &gotgbot.InlineKeyboardMarkup{
					InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
						{
							{
								Text: "Join Now",
								Url:  "https://t.me/ttonstudents_bot?start=12345678",
							},
						},
					},
				},
			},
		}

		_, err := b.AnswerInlineQuery(ctx.InlineQuery.Id, results, &gotgbot.AnswerInlineQueryOpts{
			CacheTime: 0,
		})
		if err != nil {
			return fmt.Errorf("failed to answer inline query: %w", err)
		}
	}

	return nil
}
