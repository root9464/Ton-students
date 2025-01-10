package jwt_dto

type UserData struct {
	ID       int64  `json:"id" validate:"required"`
	Username string `json:"username" validate:"required"`
}

type UserDataSаlte struct {
	Hash         string `json:"hash" validate:"required"`
	ChatInstance int64  `json:"chat_instance"`
}
