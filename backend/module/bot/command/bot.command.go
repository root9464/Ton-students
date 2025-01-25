package bot_command

import (
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
		ReplyMarkup:keyboards.Keyboard, 
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
	
	_, err := b.SendMessage(cm.config.AdminId,
		fmt.Sprintf(
			"📩 <b>Новый запрос от пользователя</b>\n\n<b>Пользователь:</b> @%s\n<b>ID:</b> <code>%d</code>\n\n<b>Вопрос:</b>\n%s",
			ctx.EffectiveUser.Username, userID, question),
		&gotgbot.SendMessageOpts{ParseMode: "HTML"})
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

func (cm *botCommand) SupportReply(b *gotgbot.Bot, ctx *ext.Context) error {
	callbackData := ctx.CallbackQuery.Data
	cm.logger.Infof("SupportReply called with callbackData: %v", callbackData)

	userIDStr := strings.TrimPrefix(callbackData, "reply_")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		cm.logger.Errorf("Error parsing userID from callbackData: %v", err)
		return err
	}

	replyState.mu.Lock()
	replyState.active[ctx.EffectiveUser.Id] = userID
	replyState.mu.Unlock()

	_, err = b.SendMessage(ctx.EffectiveUser.Id,
		"✍️ <b>Введите ваше сообщение для пользователя:</b>",
		&gotgbot.SendMessageOpts{ParseMode: "HTML"})
	return err
}

func (cm *botCommand) SendAdminResponse(b *gotgbot.Bot, ctx *ext.Context) error {
	adminID := ctx.EffectiveUser.Id
	messageText := ctx.EffectiveMessage.Text
	fmt.Println(replyState.active)

	cm.logger.Infof("SendAdminResponse called with adminID: %v, messageText: %v", adminID, messageText)

	replyState.mu.Lock()
	userID, ok := replyState.active[adminID]
	if !ok {
		replyState.mu.Unlock()
		_, err := b.SendMessage(adminID,
			"⚠️ <b>Нет активного запроса для ответа.</b>\n\nСначала нажмите кнопку \"Ответить\".",
			&gotgbot.SendMessageOpts{ParseMode: "HTML"})
		return err
	}
	delete(replyState.active, adminID)
	replyState.mu.Unlock()

	_, err := b.SendMessage(userID,
		fmt.Sprintf(
			"📬 <b>Ответ от администратора:</b>\n\n%s",
			messageText),
		&gotgbot.SendMessageOpts{ParseMode: "HTML"})
	if err != nil {
		cm.logger.Errorf("Error sending admin response to user: %v", err)
		return err
	}

	_, err = b.SendMessage(adminID,
		"✅ <b>Ваш ответ был успешно отправлен пользователю.</b>",
		&gotgbot.SendMessageOpts{ParseMode: "HTML"})
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
