package handler

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/sirupsen/logrus"
)

// lastMessageID хранит ID последнего сообщения пользователя
var lastMessageID = make(map[int64]int)

// HandleUserMessage обрабатывает текстовые сообщения и удаляет предыдущее
func HandleUserMessage(b *gotgbot.Bot, ctx *ext.Context) error {
	log := logrus.New()

	userID := ctx.EffectiveUser.Id
	chatID := ctx.EffectiveChat.Id
	messageID := ctx.EffectiveMessage.MessageId
	

	// Удаляем предыдущее сообщение, если оно есть
	if lastID, exists := lastMessageID[userID]; exists {
		_, err := b.DeleteMessage(chatID, int64(lastID), nil)
		if err != nil {
			log.Warnf("Не удалось удалить сообщение %d: %v", lastID, err)
		}
	}

	// Сохраняем ID текущего сообщения
	lastMessageID[userID] = int(messageID)

	return nil
}
