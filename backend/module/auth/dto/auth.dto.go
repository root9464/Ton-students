package auth_dto

type AutorizeDto struct {
	InitDataRaw string `json:"init-data-raw" validate:"required"`
}
