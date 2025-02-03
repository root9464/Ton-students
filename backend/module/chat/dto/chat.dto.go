package chat_dto

type CreateChatType struct {
	Users []int64 `json:"users" validate:"required"`
}
