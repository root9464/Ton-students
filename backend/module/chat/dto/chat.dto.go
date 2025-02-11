package chat_dto

type CreateChatType struct {
	Users []int64 `json:"users" validate:"required"`
}

type CreateOrLoad struct {
	UserID    int64  `json:"user_id" validate:"required"`
	ServiceID string `json:"service_id" validate:"required"`
}
