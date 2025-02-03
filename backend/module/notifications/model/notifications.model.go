package notifications_model

import "time"

type NotificationType string

const (
	InfoNotification    NotificationType = "info"
	EventNotification   NotificationType = "event"
	InviteNotification  NotificationType = "invite"
	CommentNotification NotificationType = "comment"
	MessageNotification NotificationType = "message"
	LikeNotification    NotificationType = "like"
	DislikeNotification NotificationType = "dislike"
)

type Notification struct {
	ID         string           `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	SenderID   int64            `gorm:"not null;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"sender"`
	ReceiverID int64            `gorm:"not null" json:"receiver"`
	Content    string           `gorm:"type:text" json:"content"`
	Type       NotificationType `gorm:"type:notification_type;not null" json:"type"`
	CreatedAt  time.Time        `gorm:"autoCreateTime" json:"createdAt"`
}
