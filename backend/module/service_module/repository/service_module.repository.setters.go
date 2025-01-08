package serv_repository

import (
	"context"

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
