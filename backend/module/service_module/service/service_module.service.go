package serv_service

import (
	"context"

	"github.com/go-playground/validator/v10"
	serv_dto "github.com/root9464/Ton-students/module/service_module/dto"
	serv_repository "github.com/root9464/Ton-students/module/service_module/repository"
	user_repository "github.com/root9464/Ton-students/module/user/repository"
	"github.com/root9464/Ton-students/shared/logger"
	"gorm.io/gorm"
)

var _ IServiceModuleService = (*serviceModuleService)(nil)

type IServiceModuleService interface {
	CreateService(ctx context.Context, dto *serv_dto.ServiceType) error
	UpdateInformation(ctx context.Context, dto *serv_dto.UpdateServiceType) error

	GetServiceById(ctx context.Context, id string) (*serv_dto.GetServicesType, error)
	GetShortServices(ctx context.Context, page, size int) (*[]serv_dto.FeedServiceType, error)
}

type serviceModuleService struct {
	logger    *logger.Logger
	validator *validator.Validate

	db       *gorm.DB
	repo     serv_repository.IServiceModuleRepository
	userRepo user_repository.IUserRepository
}

func NewServiceModuleService(
	logger *logger.Logger, validator *validator.Validate,
	repo serv_repository.IServiceModuleRepository,
	userRepo user_repository.IUserRepository,
	db *gorm.DB,
) *serviceModuleService {
	return &serviceModuleService{
		logger:    logger,
		validator: validator,
		repo:      repo,
		userRepo:  userRepo,
		db:        db,
	}
}
