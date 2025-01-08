package serv_repository

import (
	"context"

	serv_model "github.com/root9464/Ton-students/module/service_module/model"
	"github.com/root9464/Ton-students/shared/logger"
	"gorm.io/gorm"
)

type IServiceModuleRepository interface {
	CreateService(ctx context.Context, service *serv_model.Service) error
	GetServiceById(ctx context.Context, id string) (*serv_model.Service, error)
}

type serviceRepository struct {
	db     *gorm.DB
	logger *logger.Logger
}

func NewServiceModuleRepository(db *gorm.DB, logger *logger.Logger) *serviceRepository {
	return &serviceRepository{
		db:     db,
		logger: logger,
	}
}
