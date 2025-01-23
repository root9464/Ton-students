package serv_repository

import (
	"context"

	serv_model "github.com/root9464/Ton-students/module/service_module/model"
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
