package serv_repository

import (
	"context"

	serv_model "github.com/root9464/Ton-students/module/service_module/model"
	user_model "github.com/root9464/Ton-students/module/user/model"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

func (r *serviceRepository) GetServiceById(ctx context.Context, id string) (*serv_model.Service, error) {
	r.logger.Info("Getting service by ID...")

	service := new(serv_model.Service)

	if err := r.db.WithContext(ctx).Preload("Infos").Preload("Tags").Preload("Settings").Where("id = ?", id).First(&service).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			r.logger.Infof("Service with ID %s not found", id)
			return nil, nil
		}
		r.logger.Errorf("Error getting service by ID: %v", err)
		return nil, err
	}

	r.logger.Infof("Service with ID %s retrieved successfully", id)
	return service, nil
}

func (r *serviceRepository) GetAllServices(ctx context.Context) (*[]ServiceWithUser, error) {
	r.logger.Info("Getting all services...")

	services := new([]serv_model.Service)
	if err := r.db.WithContext(ctx).Preload("Infos").Preload("Tags").Preload("Settings").Find(&services).Error; err != nil {
		r.logger.Errorf("Error getting all services: %v", err)
		return nil, err
	}

	users := new([]user_model.User)
	if err := r.db.WithContext(ctx).Find(&users).Error; err != nil {
		r.logger.Errorf("Error getting all users: %v", err)
		return nil, err
	}

	userMap := lo.KeyBy(*users, func(u user_model.User) int64 {
		return u.ID
	})

	servicesWithUsers := lo.Map(*services, func(service serv_model.Service, _ int) ServiceWithUser {
		user, userExists := userMap[service.UserId]
		if !userExists {
			r.logger.Errorf("User not found for service ID: %s", service.ID)
		}

		return ServiceWithUser{
			UserID:    user.ID,
			Username:  user.Username,
			Hash:      user.Hash,
			ServiceID: service.ID,
			Price:     service.Price,
			Infos:     service.Infos,
			Tags:      service.Tags,
			Settings:  service.Settings,
		}
	})

	r.logger.Info("All services retrieved successfully")
	return &servicesWithUsers, nil
}
