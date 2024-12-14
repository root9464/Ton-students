package handler

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

const requiredChannel = -1002455475992 // Замените на ваш канал

func Start(b *gotgbot.Bot, ctx *ext.Context) error {
	userID := ctx.EffectiveUser.Id

	// Проверяем, подписан ли пользователь на канал
	member, err := b.GetChatMember(requiredChannel, userID, nil)
	if err != nil {
		// Обработка ошибок, если не удается получить информацию о членстве
		_, _ = ctx.EffectiveMessage.Reply(b, "Ошибка при проверке подписки. Попробуйте позже.", nil)
		return err
	}

	//Если пользователь не подписан на канал
	if member.GetStatus() != "member" && member.GetStatus() != "administrator" && member.GetStatus() != "creator" {
		// Сообщаем пользователю, что нужно подписаться
		_, _ = ctx.EffectiveMessage.Chat.SendMessage(b, "Чтобы использовать этот сервис, пожалуйста, подпишитесь на канал [здесь](https://t.me/+M3Gb_96fCzNhMWFi)", nil)
		return nil
	}

	// Если пользователь подписан, отправляем приветственное сообщение
	_, err = b.SendMessage(ctx.EffectiveChat.Id, "Hello! How can I help you?", nil)
	if err != nil {
		// Логирование ошибки, если сообщение не отправлено
		return err
	}

	return nil
}
