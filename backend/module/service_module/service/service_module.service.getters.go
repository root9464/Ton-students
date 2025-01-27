package serv_service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	serv_dto "github.com/root9464/Ton-students/module/service_module/dto"
	serv_model "github.com/root9464/Ton-students/module/service_module/model"
	"github.com/root9464/Ton-students/shared/utils"
	"github.com/samber/lo"
)

func (s *serviceModuleService) GetServiceById(ctx context.Context, id string) (*serv_dto.GetServicesType, error) {

	serviceInDB, err := s.repo.GetServiceById(ctx, id)
	if err != nil || serviceInDB == nil {
		return nil, &fiber.Error{
			Code:    404,
			Message: "Service not found",
		}
	}

	service, err := utils.ConvertDtoToEntity[serv_dto.GetServicesType](serviceInDB)
	if err != nil {
		return nil, &fiber.Error{
			Code:    500,
			Message: err.Error()}
	}

	service.Username = utils.GetVisibleName(&serviceInDB.User)

	return service, nil
}

func (s *serviceModuleService) GetShortServices(ctx context.Context, page int, size int) (*[]serv_dto.FeedServiceType, error) {
	s.logger.Info("Getting creator services...")

	services, err := s.repo.UserServices(ctx, page, size)
	if err != nil {
		return nil, &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	s.logger.Infof("Got services: %+v", services)

	if services == nil || len(*services) == 0 {
		s.logger.Infof("No services found for page %d and size %d", page, size)
		return nil, nil
	}

	shortServices := lo.Map(*services, func(service serv_model.Service, _ int) serv_dto.FeedServiceType {

		infos := lo.Map(service.Infos, func(info serv_model.ServiceInfo, _ int) serv_dto.InfosType {
			return serv_dto.InfosType{
				Title:   info.Title,
				Content: info.Content,
			}
		})

		tags := new([]serv_dto.TagsType)
		if len(service.Tags) > 0 {
			mappedTags := lo.Map(service.Tags, func(tag serv_model.Tags, _ int) serv_dto.TagsType {
				return serv_dto.TagsType{
					ServiceId: tag.ServiceId,
					Name:      tag.Name,
				}
			})
			tags = &mappedTags
		}

		settings, _ := utils.ConvertDtoToEntity[serv_dto.SettingsType](service.Settings)
		visibleName := utils.GetVisibleName(&service.User)
		result := serv_dto.FeedServiceType{
			ID:       service.ID,
			UserID:   service.UserID,
			Username: &visibleName,
			Price:    service.Price,
			Infos:    &infos,
			Settings: settings,
		}

		if tags != nil && len(*tags) > 0 {
			result.Tags = tags
		}

		return result
	})

	return &shortServices, nil
}
