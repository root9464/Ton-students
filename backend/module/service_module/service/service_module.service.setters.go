package serv_service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	serv_dto "github.com/root9464/Ton-students/module/service_module/dto"
	serv_model "github.com/root9464/Ton-students/module/service_module/model"
	"github.com/root9464/Ton-students/shared/utils"
)

func (s *serviceModuleService) CreateService(ctx context.Context, dto *serv_dto.ServiceType) error {
	s.logger.Infof("dto received: %+v", dto)
	if err := s.validator.Struct(dto); err != nil {
		s.logger.Warnf("validate error: %s", err.Error())
		return &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	s.logger.Infof("Validating dto success : %+v", dto)
	if dto.Settings.ButtonText == nil && dto.Settings.IsAdditionalButton {
		return &fiber.Error{
			Code:    422,
			Message: "buttonText must be filled if isAdditionalButton is true",
		}
	}

	s.logger.Infof("converting dto to entity: %+v", dto)
	newServicem, err := utils.ConvertDtoToEntity[serv_model.Service](dto)
	if err != nil {
		s.logger.Warnf("convert dto to entity error: %s", err.Error())
		return &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	if err := s.repo.CreateService(ctx, newServicem); err != nil {
		s.logger.Warnf("create service error: %s", err.Error())
		return &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	return nil
}
