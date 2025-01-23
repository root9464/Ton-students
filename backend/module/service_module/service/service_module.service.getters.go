package serv_service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	serv_dto "github.com/root9464/Ton-students/module/service_module/dto"
	serv_model "github.com/root9464/Ton-students/module/service_module/model"
	"github.com/root9464/Ton-students/shared/utils"
	"github.com/samber/lo"
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

	if len(*service.Tags) == 0 {
		service.Tags = nil
	}

	return service, nil
}

func (s *serviceModuleService) GetShortServices(ctx context.Context) (*[]serv_dto.ShortServiceType, error) {
	s.logger.Info("Getting creator services...")

	user, err := s.userRepo.UserServices(ctx)
	if err != nil {
		return nil, &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	s.logger.Infof("Got creator services: %+v", user)

	// Проверяем, что Services не nil и не пустой
	if user == nil || user.Services == nil || len(user.Services) == 0 {
		s.logger.Info("No services found for user")
		return nil, nil
	}

	shortServices := lo.Map(user.Services, func(service serv_model.Service, _ int) serv_dto.ShortServiceType {
		infos := lo.Map(service.Infos, func(info serv_model.ServiceInfo, _ int) serv_dto.InfosType {
			return serv_dto.InfosType{
				Title:   info.Title,
				Content: info.Content,
			}
		})

		tags := new([]serv_dto.TagsType)
		if service.Tags != nil && len(*service.Tags) > 0 {
			mappedTags := lo.Map(*service.Tags, func(tag serv_model.Tags, _ int) serv_dto.TagsType {
				return serv_dto.TagsType{
					ServiceId: tag.ServiceId,
					Name:      tag.Name,
				}
			})
			tags = &mappedTags
		}

		settings, _ := utils.ConvertDtoToEntity[serv_dto.SettingsType](service.Settings)

		visibleName := utils.GetVisibleName(user)

		result := serv_dto.ShortServiceType{
			ID:       service.ID,
			UserID:   service.UserId,
			Username: &visibleName,
			Price:    service.Price,
			Infos:    infos[0],
			Settings: settings,
		}

		if tags != nil && len(*tags) > 0 {
			result.Tags = tags
		}

		return result
	})

	s.logger.Infof("Short creator services retrieved: count = %d", len(shortServices))

	return &shortServices, nil
}
