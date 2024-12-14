package handler

import (
	"fmt"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/sirupsen/logrus"
)

var userToAdmin = make(map[int64]int64)


func SupportStart(b *gotgbot.Bot, ctx *ext.Context, adminID int64, log *logrus.Logger) error {
	userID := ctx.EffectiveUser.Id

	// Проверяем, отправил ли пользователь текстовый вопрос
	args := ctx.Args()
	if len(args) == 0 {
		// Если вопрос не указан, отправляем инструкцию
		_, err := ctx.EffectiveMessage.Reply(b, "Пожалуйста, укажите ваш вопрос. Пример:\n/start Как зарегистрироваться?", nil)
		if err != nil {
			log.Error(err)
			return err
		}
		return nil
	}

	// Объединяем все слова после /start в вопрос
	question := strings.Join(args, " ")

	// Отправляем уведомление админу
	_, err := b.SendMessage(adminID, fmt.Sprintf("Новый запрос от пользователя @%s (ID: %d):\n%s", ctx.EffectiveUser.Username, userID, question), nil)
	if err != nil {
		log.Error(err)
		return err
	}

	// Подтверждаем пользователю
	_, err = b.SendMessage(userID, "Ваш запрос был отправлен в поддержку. Ожидайте ответа от администратора.", nil)
	if err != nil {
		log.Error(err)
		return err
	}

	// Сохраняем связь пользователя с админом
	userToAdmin[userID] = adminID

	return nil
}
