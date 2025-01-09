package serv_repository

import (
	"context"
	"fmt"

	serv_model "github.com/root9464/Ton-students/module/service_module/model"
)

func (r *serviceRepository) CreateService(ctx context.Context, service *serv_model.Service) error {
	r.logger.Info("Creating service...")
	if err := r.db.WithContext(ctx).Create(&service).Error; err != nil {
		r.logger.Errorf("Error creating service: %v", err)
		return err
	}
	r.logger.Info("Service created successfully")
	return nil
}

func (r *serviceRepository) UpdateService(ctx context.Context, service *serv_model.Service) error {
	r.logger.Info("Updating service...")
	result := r.db.WithContext(ctx).Model(&serv_model.Service{}).Where("id = ?", service.ID).Updates(service)
	if result.Error != nil {
		r.logger.Errorf("Error updating service: %v", result.Error)
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("service with ID '%s' does not exist", service.ID)
	}
	r.logger.Info("Service updated successfully")
	return nil
}
