package service_dto

type CreateServiceDto struct {
	UserID      int64                  `json:"user_id" validate:"required"`
	Title       string                 `json:"title" validate:"required"`
	Description map[string]interface{} `json:"description"`
	Tags        []string               `json:"tags"`
	Price       int16                  `json:"price" validate:"required"`
}
