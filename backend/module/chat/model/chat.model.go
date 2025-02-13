package chat_model

import (
	"time"

	user_model "github.com/root9464/Ton-students/module/user/model"
)

type Chat struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	ServiceID string    `gorm:"not null"`
	Key       string    `json:"key"`
	CreatedAt time.Time `json:"createdAt"`

	Users    []user_model.User `gorm:"many2many:chat_users;"`
	Messages []Message         `gorm:"foreignKey:ChatID" json:"messages"`
}

type Message struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	ChatID    string    `gorm:"not null;index" json:"chatId"`
	SenderID  int64     `gorm:"not null;index" json:"senderId"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
}
