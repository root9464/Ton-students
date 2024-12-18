package user_dto

type CreateUserDto struct {
	InitDataRaw string `json:"init-data-raw" validate:"required"`
}
