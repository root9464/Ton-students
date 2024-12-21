package service_repository

import (
	"context"

	"github.com/root9464/Ton-students/ent"
	"github.com/root9464/Ton-students/shared/logger"
)

type IServiceRepository interface {
	Create(ctx context.Context, dto *ent.Service) (*ent.Service, error)
}

type serviceRepository struct {
	db     *ent.Client
	logger *logger.Logger
}

func NewServiceRepository(
	db *ent.Client,
	logger *logger.Logger,
) *serviceRepository {
	return &serviceRepository{
		db:     db,
		logger: logger,
	}
}
