package common_model

type ChatUser struct {
	ChatID string `gorm:"primaryKey" json:"chat_id"`
	UserID int64  `gorm:"primaryKey" json:"user_id"`
}
