package serv_service

import (
	"context"
	"strings"
	"sync"

	"github.com/gofiber/fiber/v2"
	serv_dto "github.com/root9464/Ton-students/module/service_module/dto"
	serv_model "github.com/root9464/Ton-students/module/service_module/model"
	"github.com/root9464/Ton-students/shared/utils"
	"github.com/samber/lo"
)

func (s *serviceModuleService) GetServiceById(ctx context.Context, id string) (*serv_dto.GetServicesType, error) {
	s.logger.Infof("Getting service by ID: %s", id)
	serviceInDB, err := s.repo.GetServiceById(ctx, id)
	if err != nil || serviceInDB == nil {
		return nil, &fiber.Error{
			Code:    404,
			Message: "Service not found",
		}
	}

	s.logger.Infof("Service with ID %v found", serviceInDB)
	service, err := utils.ConvertDtoToEntity[serv_dto.GetServicesType](serviceInDB)
	if err != nil {
		return nil, &fiber.Error{
			Code:    500,
			Message: err.Error()}
	}

	service.Nickname = serviceInDB.User.Nickname

	s.logger.Infof("Service with ID %s retrieved successfully", id)

	return service, nil
}

func (s *serviceModuleService) GetShortServices(ctx context.Context, page int, size int) (*[]serv_dto.FeedServiceType, int64, int64, error) {
	s.logger.Info("Getting creator services...")

	var (
		wg          sync.WaitGroup
		total       int64
		services    []serv_model.Service
		totalErr    error
		servicesErr error
	)

	wg.Add(2)

	go func() {
		defer wg.Done()
		res, err := s.repo.TotalServices(ctx)
		total, totalErr = res, err
	}()

	go func() {
		defer wg.Done()
		res, err := s.repo.UserServices(ctx, page, size)
		if err != nil {
			servicesErr = err
			return
		}
		services = *res
	}()

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
	case <-ctx.Done():
		return nil, 0, 0, ctx.Err()
	}

	if totalErr != nil || servicesErr != nil {
		errMsg := ""
		if totalErr != nil {
			errMsg += "Total error: " + totalErr.Error()
		}
		if servicesErr != nil {
			errMsg += " Services error: " + servicesErr.Error()
		}
		return nil, 0, 0, &fiber.Error{Code: 500, Message: strings.TrimSpace(errMsg)}
	}

	if len(services) == 0 {
		s.logger.Infof("No services found for page %d and size %d", page, size)
		return nil, total, 0, nil
	}

	s.logger.Infof("Getting %d services for page %d and size %d", len(services), page, size)
	shortServices := lo.Map(services, func(service serv_model.Service, _ int) serv_dto.FeedServiceType {
		infos := lo.Map(service.Infos, func(info serv_model.ServiceInfo, _ int) serv_dto.InfosType {
			return serv_dto.InfosType{Title: info.Title, Content: info.Content}
		})

		tags := make([]serv_dto.TagsType, 0)
		if len(service.Tags) > 0 {
			tags = lo.Map(service.Tags, func(tag serv_model.Tags, _ int) serv_dto.TagsType {
				return serv_dto.TagsType{ServiceId: tag.ServiceId, Name: tag.Name}
			})
		}

		nilIfEmpty := func(s []serv_dto.TagsType) *[]serv_dto.TagsType {
			if len(s) == 0 {
				return nil
			}
			return &s
		}

		settings, _ := utils.ConvertDtoToEntity[serv_dto.SettingsType](service.Settings)

		return serv_dto.FeedServiceType{
			ID:       service.ID,
			UserID:   service.UserID,
			Nickname: service.User.Nickname,
			Price:    service.Price,
			Infos:    &infos,
			Settings: settings,
			Tags:     nilIfEmpty(tags),
		}
	})

	return &shortServices, total, (total / int64(size)), nil
}
