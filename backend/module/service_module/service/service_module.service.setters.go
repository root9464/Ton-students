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

	s.logger.Infof("create service success: %+v", newServicem)
	return nil
}

func (s *serviceModuleService) UpdateInformation(ctx context.Context, dto *serv_dto.UpdateServiceType) error {
	s.logger.Infof("dto received: %+v", dto)
	if err := s.validator.Struct(dto); err != nil {
		s.logger.Warnf("validate error: %s", err.Error())
		return &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}
	s.logger.Infof("Validating dto success : %+v", dto)

	if dto.Price == nil && len(dto.Infos) == 0 && len(dto.Tags) == 0 {
		return &fiber.Error{
			Code:    400,
			Message: "At least one field for update must be provided",
		}
	}

	if dto.Price != nil {
		err := s.repo.UpdateServicePrice(ctx, dto.ID, *dto.Price)
		if err != nil {
			return err
		}
	}

	if len(dto.Infos) > 0 {
		for _, info := range dto.Infos {
			s.logger.Infof("converting dto to entity: %+v", info)
			serviceInfo, err := utils.ConvertDtoToEntity[serv_model.ServiceInfo](info)
			if err != nil {
				return &fiber.Error{
					Code:    500,
					Message: err.Error(),
				}
			}
			s.logger.Infof("converting dto to entity success: %+v", serviceInfo)

			s.logger.Infof("update service info: %+v", serviceInfo)
			err = s.repo.UpdateServiceInfo(ctx, serviceInfo)
			if err != nil {
				return err
			}
		}
	}

	s.logger.Infof("Service updated successfully")
	return nil
}
