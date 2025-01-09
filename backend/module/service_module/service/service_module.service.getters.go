package serv_service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	serv_dto "github.com/root9464/Ton-students/module/service_module/dto"
	"github.com/root9464/Ton-students/shared/utils"
)

func (s *serviceModuleService) GetServiceById(ctx context.Context, id string) (*serv_dto.ServiceType, error) {

	serviceInDB, err := s.repo.GetServiceById(ctx, id)
	if err != nil || serviceInDB == nil {
		return nil, &fiber.Error{
			Code:    404,
			Message: "Service not found",
		}
	}

	service, err := utils.ConvertDtoToEntity[serv_dto.ServiceType](serviceInDB)
	if err != nil {
		return nil, &fiber.Error{
			Code:    500,
			Message: err.Error()}
	}

	return service, nil
}
