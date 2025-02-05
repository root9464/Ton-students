package serv_service

import (
	"context"
	"fmt"
	"sync"

	"github.com/gofiber/fiber/v2"
	serv_dto "github.com/root9464/Ton-students/module/service_module/dto"
	serv_model "github.com/root9464/Ton-students/module/service_module/model"
	"github.com/root9464/Ton-students/shared/utils"
)

var (
	wg sync.WaitGroup
	mu sync.Mutex
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
	s.logger.Infof("DTO received: %+v", dto)

	if err := s.validator.Struct(dto); err != nil {
		s.logger.Warnf("Validation error: %s", err.Error())
		return &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}
	s.logger.Infof("DTO validation successful: %+v", dto)

	if dto.Price == nil && len(dto.Infos) == 0 && len(dto.Tags) == 0 {
		return &fiber.Error{
			Code:    400,
			Message: "At least one field for update must be provided",
		}
	}

	var (
		updatedFields []string
		errors        []error
	)

	processError := func(err error) {
		mu.Lock()
		errors = append(errors, err)
		mu.Unlock()
	}

	s.logger.Infof("Updating service price: %+v", dto.Price)
	if dto.Price != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			taskCtx, cancel := context.WithCancel(ctx)
			defer cancel()

			if err := s.repo.UpdateServicePrice(taskCtx, dto.ID, *dto.Price); err != nil {
				s.logger.Errorf("Failed to update service price for service ID: %s. Error: %s", dto.ID, err.Error())
				processError(err)
				return
			}

			s.logger.Infof("Price updated successfully for service ID: %s", dto.ID)
			updatedFields = append(updatedFields, "price")
		}()
	}

	s.logger.Infof("Updating service information: %+v", dto.Infos)
	if len(dto.Infos) > 0 {
		for _, info := range dto.Infos {
			wg.Add(1)
			go func(info serv_dto.UpdateInfoType) {
				defer wg.Done()
				taskCtx, cancel := context.WithCancel(ctx)
				defer cancel()

				if info.ID == "" {
					err := fmt.Errorf("empty ID for service info")
					s.logger.Errorf("Failed to update service info: %s", err.Error())
					processError(err)
					return
				}

				serviceInfo, err := utils.ConvertDtoToEntity[serv_model.ServiceInfo](info)
				if err != nil {
					s.logger.Errorf("Failed to convert DTO to entity for service ID: %s. Error: %s", dto.ID, err.Error())
					processError(err)
					return
				}

				if err := s.repo.UpdateServiceInfo(taskCtx, serviceInfo); err != nil {
					s.logger.Errorf("Failed to update service info for service ID: %s. Error: %s", dto.ID, err.Error())
					processError(err)
					return
				}

				s.logger.Infof("Service info updated successfully for service ID: %s", dto.ID)
				updatedFields = append(updatedFields, "infos")
			}(info)
		}
	}

	s.logger.Infof("Updating service tags: %+v", dto.Tags)
	if len(dto.Tags) > 0 {
		for _, tag := range dto.Tags {
			wg.Add(1)
			go func(tag serv_dto.UpdateTagsType) {
				defer wg.Done()
				taskCtx, cancel := context.WithCancel(ctx)
				defer cancel()

				if tag.ID == "" {
					err := fmt.Errorf("empty ID for service tag")
					s.logger.Errorf("Failed to update service tag: %s", err.Error())
					processError(err)
					return
				}

				serviceTag, err := utils.ConvertDtoToEntity[serv_model.Tags](tag)
				if err != nil {
					s.logger.Errorf("Failed to convert DTO to entity for service ID: %s. Error: %s", dto.ID, err.Error())
					processError(err)
					return
				}

				if err := s.repo.UpdateServiceTag(taskCtx, serviceTag); err != nil {
					s.logger.Errorf("Failed to update service tag for service ID: %s. Error: %s", dto.ID, err.Error())
					processError(err)
					return
				}

				s.logger.Infof("Service tag updated successfully for service ID: %s", dto.ID)
				updatedFields = append(updatedFields, "tags")
			}(tag)
		}
	}
	wg.Wait()

	if len(updatedFields) > 0 {
		s.logger.Infof("Successfully updated fields: %v for service ID: %s", updatedFields, dto.ID)
	}

	if len(errors) > 0 {
		s.logger.Errorf("Errors occurred during update: %v", errors)
		return &fiber.Error{
			Code:    207,
			Message: "Some updates failed. Please try again or contact support.",
		}
	}

	s.logger.Infof("All updates completed successfully for service ID: %s", dto.ID)
	return nil
}
