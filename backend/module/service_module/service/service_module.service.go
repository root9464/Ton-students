package serv_service

import (
	"context"

	"github.com/go-playground/validator/v10"
	serv_dto "github.com/root9464/Ton-students/module/service_module/dto"
	serv_repository "github.com/root9464/Ton-students/module/service_module/repository"
	"github.com/root9464/Ton-students/shared/logger"
	"gorm.io/gorm"
)

var _ IServiceModuleService = (*serviceModuleService)(nil)

type IServiceModuleService interface {
	CreateService(ctx context.Context, dto *serv_dto.ServiceType) error
	UpdateInformation(ctx context.Context, dto *serv_dto.UpdateServiceType) error

	GetServiceById(ctx context.Context, id string) (*serv_dto.ServiceType, error)
	GetShortServices(ctx context.Context) (*[]serv_dto.ShortServiceType, error)
}

type serviceModuleService struct {
	logger    *logger.Logger
	validator *validator.Validate

	db   *gorm.DB
	repo serv_repository.IServiceModuleRepository
}

func NewServiceModuleService(logger *logger.Logger, validator *validator.Validate, repo serv_repository.IServiceModuleRepository, db *gorm.DB) *serviceModuleService {
	return &serviceModuleService{logger: logger, validator: validator, repo: repo, db: db}
}
