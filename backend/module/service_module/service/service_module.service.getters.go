package serv_service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	serv_dto "github.com/root9464/Ton-students/module/service_module/dto"
	serv_model "github.com/root9464/Ton-students/module/service_module/model"
	user_model "github.com/root9464/Ton-students/module/user/model"
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

func (s *serviceModuleService) GetShortServices(ctx context.Context, page int, size int) (*[]serv_dto.FeedServiceType, error) {
	s.logger.Info("Getting creator services...")

	users, err := s.userRepo.UserServices(ctx)
	if err != nil {
		return nil, &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	s.logger.Infof("Got creator services: %+v", users)

	if users == nil || len(*users) == 0 {
		s.logger.Infof("No services found for users %+v", users)
		return nil, nil
	}

	start := (page - 1) * size
	end := start + size

	shortServices := lo.FlatMap(*users, func(user user_model.User, _ int) []serv_dto.FeedServiceType {
		visibleName := utils.GetVisibleName(&user)

		return lo.Map(user.Services, func(service serv_model.Service, _ int) serv_dto.FeedServiceType {
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
	})

	if start > len(shortServices) {
		return nil, nil
	}

	if end > len(shortServices) {
		end = len(shortServices)
	}

	paginatedServices := shortServices[start:end]

	s.logger.Infof("Short creator services retrieved: count = %d", len(paginatedServices))

	return &paginatedServices, nil
}
