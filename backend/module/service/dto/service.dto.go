package service_dto

type CreateServiceDto struct {
	Name        string                 `json:"name" validate:"required"`
	Title       string                 `json:"title" validate:"required"`
	Description map[string]interface{} `json:"description"`
	Tags        []string               `json:"tags"`
	Price       int16                  `json:"price" validate:"required"`
}
