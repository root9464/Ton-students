package chat_model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Chat struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Key       string    `json:"key"`
	CreatedAt time.Time `json:"createdAt"`
	Messages  []Message `gorm:"foreignKey:ChatID" json:"messages"`
}

func (c *Chat) BeforeCreate(tx *gorm.DB) error {
	c.ID = uuid.New().String()
	return nil
}

type Message struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	ChatID    string    `gorm:"not null;index" json:"chatId"`
	SenderID  int64     `gorm:"not null;index" json:"senderId"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
}
