package common_model

import (
	chat_model "github.com/root9464/Ton-students/module/chat/model"
	user_model "github.com/root9464/Ton-students/module/user/model"
)

type ChatUser struct {
	ChatID string `gorm:"primaryKey" json:"chat_id"`
	UserID int64  `gorm:"primaryKey" json:"user_id"`

	Chat *chat_model.Chat `gorm:"foreignKey:ChatID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"chat"`
	User *user_model.User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user"`
}
