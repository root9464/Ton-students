package bot_dto

type Payment struct {
	UserID int64 `json:"userId" validate:"required"`
}
