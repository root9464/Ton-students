package bot_service

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/sirupsen/logrus"
)

func (c *BotService) SupportStart(b *gotgbot.Bot, ctx *ext.Context) error {
	c.logger.WithFields(logrus.Fields{
		"userID": ctx.EffectiveUser.Id,
	}).Info("SupportStart called")

	userID := ctx.EffectiveUser.Id
	args := ctx.Args()

	if len(args) == 0 {
		_, err := ctx.EffectiveMessage.Reply(b,
			"❓ <b>Введите ваш вопрос</b>\n\nПример:\n<code>/support Как зарегистрироваться?</code>",
			&gotgbot.SendMessageOpts{ParseMode: "HTML"})
		return err
	}

	question := strings.Join(args, " ")
	c.logger.WithFields(logrus.Fields{
		"userID":   userID,
		"question": question,
	}).Info("Received support question")

	_, err := b.SendMessage(c.config.AdminId,
		fmt.Sprintf(
			"📩 <b>Новый запрос от пользователя</b>\n\n<b>Пользователь:</b> @%s\n<b>ID:</b> <code>%d</code>\n\n<b>Вопрос:</b>\n%s",
			ctx.EffectiveUser.Username, userID, question),
		&gotgbot.SendMessageOpts{ParseMode: "HTML"})
	if err != nil {
		c.logger.WithError(err).Error("Error sending support request")
		return err
	}

	_, err = b.SendMessage(userID,
		"✅ <b>Ваш запрос отправлен в поддержку</b>\n\nПожалуйста, ожидайте ответа от администратора.",
		&gotgbot.SendMessageOpts{ParseMode: "HTML"})
	if err != nil {
		c.logger.WithError(err).Error("Error sending confirmation message after support request")
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
	_, err = b.SendMessage(c.config.AdminId,
		"👤 <b>Вы можете ответить на запрос этого пользователя:</b>",
		&gotgbot.SendMessageOpts{
			ParseMode:   "HTML",
			ReplyMarkup: replyMarkup,
		})
	return err
}

// SupportReply - Логика для ответа на запросы поддержки
func (c *BotService) SupportReply(b *gotgbot.Bot, ctx *ext.Context) error {
	callbackData := ctx.CallbackQuery.Data
	c.logger.Infof("SupportReply called with callbackData:", callbackData)

	userIDStr := strings.TrimPrefix(callbackData, "reply_")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		c.logger.Errorf("Error parsing userID from callbackData:", err)
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

func (c *BotService) SendAdminResponse(b *gotgbot.Bot, ctx *ext.Context) error {
	adminID := ctx.EffectiveUser.Id
	messageText := ctx.EffectiveMessage.Text
	fmt.Println(replyState.active)

	c.logger.WithFields(logrus.Fields{
		"adminID": adminID,
		"message": messageText,
	}).Info("SendAdminResponse called")

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
		c.logger.WithError(err).Error("Error sending admin response")
		return err
	}

	_, err = b.SendMessage(adminID,
		"✅ <b>Ваш ответ был успешно отправлен пользователю.</b>",
		&gotgbot.SendMessageOpts{ParseMode: "HTML"})
	return err
}
