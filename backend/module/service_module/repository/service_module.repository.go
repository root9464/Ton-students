package serv_repository

import (
	"context"

	serv_model "github.com/root9464/Ton-students/module/service_module/model"
	user_repository "github.com/root9464/Ton-students/module/user/repository"
	"github.com/root9464/Ton-students/shared/logger"
	"gorm.io/gorm"
)

type ServiceWithUser struct {
	UserID   int64  `json:"userId"`
	Username string `json:"username"`
	Hash     string `json:"hash"`

	ServiceID string                     `json:"serviceId"`
	Price     float64                    `json:"price"`
	Infos     []serv_model.ServiceInfo   `json:"infos"`
	Tags      *[]serv_model.Tags         `json:"tags"`
	Settings  serv_model.ServiceSettings `json:"settings"`
}

type IServiceModuleRepository interface {
	CreateService(ctx context.Context, service *serv_model.Service) error
	GetServiceById(ctx context.Context, id string) (*serv_model.Service, error)
	GetAllServices(ctx context.Context) (*[]ServiceWithUser, error)
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
