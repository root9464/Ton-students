package notifications_dto

type NotificationType struct {
	Sender   string `json:"sender" validate:"required"`
	Receiver string `json:"receiver" validate:"required"`
	Content  string `json:"content"`
	Type     string `json:"type" validate:"required,oneof=info event invite comment message like dislike"`
	Created  string `json:"created" validate:"required"`
}

type CreateNotificationType struct {
	SenderID   int64  `json:"sender_id" validate:"required"`
	ReceiverID int64  `json:"receiver_id" validate:"required"`
	Content    string `json:"content"`
	Type       string `json:"type" validate:"required,oneof=info event invite comment message like dislike"`
}
