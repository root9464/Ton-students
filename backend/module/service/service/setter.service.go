package service_serv

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/root9464/Ton-students/ent"
	service_dto "github.com/root9464/Ton-students/module/service/dto"
	"github.com/root9464/Ton-students/shared/utils"
)

func (s *servService) CreateService(ctx context.Context, dto *service_dto.CreateServiceDto) (*ent.Service, error) {
	if err := s.validator.Struct(dto); err != nil {
		return nil, &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	convertEntity, err := utils.DtoToModel(dto, ent.Service{})
	if err != nil {
		return nil, &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	createdService, err := s.repo.Create(ctx, convertEntity)
	if err != nil {
		return nil, &fiber.Error{
			Code:    500,
			Message: "Error creating service: " + err.Error(),
		}
	}

	return createdService, nil
}
