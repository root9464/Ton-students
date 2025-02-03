package chat_model

import (
	"time"

	common_model "github.com/root9464/Ton-students/module/model/common"
)

type Chat struct {
	ID        string                  `gorm:"primaryKey" json:"id"`
	Members   []common_model.ChatUser `gorm:"many2many:user_chats;joinForeignKey:ChatID;joinReferences:UserID" json:"-"`
	Key       string                  `json:"key"`
	CreatedAt time.Time               `json:"createdAt"`
	Messages  []Message               `gorm:"foreignKey:ChatID" json:"messages"`
}

type Message struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	ChatID    string    `gorm:"not null;index" json:"chatId"`
	SenderID  int64     `gorm:"not null;index" json:"senderId"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
}
