package serv_repository

import (
	"context"

	serv_model "github.com/root9464/Ton-students/module/service_module/model"
	user_repository "github.com/root9464/Ton-students/module/user/repository"
	"github.com/root9464/Ton-students/shared/logger"
	"gorm.io/gorm"
)

var _ IServiceModuleRepository = (*serviceRepository)(nil)

type IServiceModuleRepository interface {
	CreateService(ctx context.Context, service *serv_model.Service) error
	UpdateServiceTag(ctx context.Context, service *serv_model.Tags) error

	UpdateServiceInfo(ctx context.Context, service *serv_model.ServiceInfo) error
	UpdateServicePrice(ctx context.Context, id string, price float64) error
	GetServiceById(ctx context.Context, id string) (*serv_model.Service, error)
	UserServices(ctx context.Context, page, size int) (*[]serv_model.Service, error)
}

type serviceRepository struct {
	db     *gorm.DB
	logger *logger.Logger

	userRepo user_repository.IUserRepository
}

func NewServiceModuleRepository(db *gorm.DB, logger *logger.Logger, userRepo user_repository.IUserRepository) *serviceRepository {
	return &serviceRepository{
		db:       db,
		logger:   logger,
		userRepo: userRepo,
	}
}
