package service_repository

import (
	"context"

	"github.com/root9464/Ton-students/ent"
	"github.com/root9464/Ton-students/shared/logger"
)

var _ IServiceRepository = (*serviceRepository)(nil)

type IServiceRepository interface {
	Create(ctx context.Context, service *ent.Service) (*ent.Service, error)
	CreateTags(ctx context.Context, tags []string) ([]*ent.Tags, error)
	CreateTag(ctx context.Context, tagName string) (*ent.Tags, error)
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
