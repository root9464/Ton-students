package bot_service

import (
	"strings"
	"sync"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/sirupsen/logrus"
)

var replyState = struct {
	mu     sync.Mutex
	active map[int64]int64
}{
	active: make(map[int64]int64),
}

// Start - Логика обработки команды /start
func (c *BotService) Start(b *gotgbot.Bot, ctx *ext.Context) error {
	//считывает аргументы команды start
	referalCode := searchArgs(ctx)
	c.logger.Info(referalCode)
	//обрабатываем id чела который дал рефералку -> если он есть в бд то добавляем в бд рефералку
	//											 -> если нету такого чела в бд ссылка не действует

	userID := ctx.EffectiveUser.Id

	c.logger.WithFields(logrus.Fields{
		"userID": userID,
	}).Info("Start command received")

	member, err := b.GetChatMember(c.config.ChannelId, userID, nil)
	if err != nil {
		c.logger.WithError(err).Error("Error checking subscription status")
		_, _ = ctx.EffectiveMessage.Reply(b, "Ошибка при проверке подписки. Попробуйте позже.", nil)
		return err
	}

	memberStatus := member.GetStatus()
	c.logger.WithFields(logrus.Fields{
		"userID":       userID,
		"memberStatus": memberStatus,
	}).Info("Subscription status checked")

	if memberStatus != "member" && memberStatus != "administrator" && memberStatus != "creator" {
		_, err := ctx.EffectiveMessage.Chat.SendMessage(b,
			"Чтобы использовать этот сервис, пожалуйста, подпишитесь на канал [здесь](https://t.me/+M3Gb_96fCzNhMWFi)",
			nil,
		)
		if err != nil {
			c.logger.WithError(err).Error("Error sending subscription requirement message")
		}
		return nil
	}

	_, err = b.SendMessage(ctx.EffectiveChat.Id, "Hello! How can I help you?", nil)
	if err != nil {
		c.logger.WithError(err).Error("Error sending welcome message")
		return err
	}

	c.logger.WithFields(logrus.Fields{
		"userID": userID,
	}).Info("Welcome message sent successfully")

	return nil
}

func searchArgs(ctx *ext.Context) string {
	commandText := ctx.EffectiveMessage.Text

	if strings.HasPrefix(commandText, "/start ") {
		return strings.TrimSpace(strings.TrimPrefix(commandText, "/start "))
	}

	return ""
}
