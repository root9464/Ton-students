package serv_repository

import (
	"context"

	serv_model "github.com/root9464/Ton-students/module/service_module/model"
	"gorm.io/gorm"
)

func (r *serviceRepository) GetServiceById(ctx context.Context, id string) (*serv_model.Service, error) {
	r.logger.Info("Getting service by ID...")

	service := new(serv_model.Service)

	if err := r.db.WithContext(ctx).Preload("User").Preload("Infos").Preload("Tags").Preload("Settings").Where("id = ?", id).First(&service).Error; err != nil {
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

func (r *serviceRepository) UserServices(ctx context.Context, page, size int) (*[]serv_model.Service, error) {
	r.logger.Info("Getting creator services...")

	services := new([]serv_model.Service)

	if err := r.db.WithContext(ctx).
		Preload("User").
		Preload("Infos").
		Preload("Tags").
		Preload("Settings").
		Offset((page - 1) * size).
		Limit(size).
		Find(&services).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			r.logger.Info("No services found")
			return nil, nil
		}
		r.logger.Errorf("Error getting services: %v", err)
		return nil, err
	}

	return services, nil
}

func (r *serviceRepository) TotalServices(ctx context.Context) (int64, error) {
	r.logger.Info("Getting total number of services...")
	count := int64(0)

	if err := r.db.WithContext(ctx).Model(&serv_model.Service{}).Count(&count).Error; err != nil {
		r.logger.Errorf("Error getting total number of services: %v", err)
		return 0, err
	}

	return count, nil
}
