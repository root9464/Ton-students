package bot_dto

type Payment struct {
	UserId int64 `json:"userId" validate:"required"`
}
