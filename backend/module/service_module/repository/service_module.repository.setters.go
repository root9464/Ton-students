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

func (r *serviceRepository) UpdateServiceTag(ctx context.Context, service *serv_model.Tags) error {
	r.logger.Info("Updating service tags...")
	result := r.db.WithContext(ctx).Model(&serv_model.Tags{}).Where("id = ?", service.ID).Updates(service)
	if result.Error != nil {
		r.logger.Errorf("Error updating service tags: %v", result.Error)
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("service with ID '%s' does not exist", service.ID)
	}
	r.logger.Info("Service tags updated successfully")
	return nil
}

func (r *serviceRepository) UpdateServiceInfo(ctx context.Context, service *serv_model.ServiceInfo) error {
	r.logger.Info("Updating service info...")
	result := r.db.WithContext(ctx).Model(&serv_model.ServiceInfo{}).Where("id = ?", service.ID).Updates(service)
	if result.Error != nil {
		r.logger.Errorf("Error updating service info: %v", result.Error)
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("service with ID '%s' does not exist", service.ID)
	}
	r.logger.Info("Service info updated successfully")
	return nil
}

func (r *serviceRepository) UpdateServicePrice(ctx context.Context, id string, price float64) error {
	r.logger.Info("Updating service price...")
	result := r.db.WithContext(ctx).Model(&serv_model.Service{}).Where("id = ?", id).Update("price", price)
	if result.Error != nil {
		r.logger.Errorf("Error updating service price: %v", result.Error)
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("service with ID '%s' does not exist", id)
	}
	r.logger.Info("Service price updated successfully")
	return nil
}
