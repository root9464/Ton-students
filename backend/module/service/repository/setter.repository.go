package service_repository

import (
	"context"

	"github.com/root9464/Ton-students/ent"
)

func (r *serviceRepository) Create(ctx context.Context, service *ent.Service) (*ent.Service, error) {
	r.logger.Info("Creating service...")

	getService, err := r.db.Service.Create().
		SetUserID(service.UserID).
		SetTitle(service.Title).
		SetDescription(service.Description).
		SetTags(service.Tags).
		SetPrice(service.Price).
		Save(ctx)

	if err != nil {
		r.logger.Errorf("Error creating service: %v", err)
		return nil, err
	}
	r.logger.Info("Service create successfully")
	return getService, nil
}
