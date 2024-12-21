package handler

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/sirupsen/logrus"
)


func Start(b *gotgbot.Bot, ctx *ext.Context,channelid int64, log *logrus.Logger) error {
	userID := ctx.EffectiveUser.Id

	log.WithFields(logrus.Fields{
		"userID": userID,
	}).Info("Start command received")

	member, err := b.GetChatMember(channelid, userID, nil)
	if err != nil {
		log.WithError(err).Error("Error checking subscription status")
		_, _ = ctx.EffectiveMessage.Reply(b, "Ошибка при проверке подписки. Попробуйте позже.", nil)
		return err
	}

	memberStatus := member.GetStatus()
	log.WithFields(logrus.Fields{
		"userID":       userID,
		"memberStatus": memberStatus,
	}).Info("Subscription status checked")

	if memberStatus != "member" && memberStatus != "administrator" && memberStatus != "creator" {
		_, err := ctx.EffectiveMessage.Chat.SendMessage(b, 
			"Чтобы использовать этот сервис, пожалуйста, подпишитесь на канал [здесь](https://t.me/+M3Gb_96fCzNhMWFi)", 
			nil,
		)
		if err != nil {
			log.WithError(err).Error("Error sending subscription requirement message")
		}
		return nil
	}

	_, err = b.SendMessage(ctx.EffectiveChat.Id, "Hello! How can I help you?", nil)
	if err != nil {
		log.WithError(err).Error("Error sending welcome message")
		return err
	}

	log.WithFields(logrus.Fields{
		"userID": userID,
	}).Info("Welcome message sent successfully")

	return nil
}
