package handler

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/sirupsen/logrus"
)

// сопоставление пользователя с адиином
var userToAdmin = make(map[int64]int64)

func SupportStart(b *gotgbot.Bot, ctx *ext.Context, adminID int64, log *logrus.Logger) error {
	userID := ctx.EffectiveUser.Id

	_, err := b.SendMessage(adminID, fmt.Sprintf("Новый запрос от пользователя @%s (ID: %d).", ctx.EffectiveUser.Username, userID), nil)
	if err != nil {
		log.Error(err)
		return err
	}

	_,err = b.SendMessage(userID,"Вы связались с поддержкой. Администратор ответит вам в ближайшее время.", nil) 
	if err != nil {
		log.Error(err)
		return err
	}

	userToAdmin[userID] = adminID
	
	return nil
}
