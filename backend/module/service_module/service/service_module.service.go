package serv_service

import (
	"context"

	"github.com/go-playground/validator/v10"
	serv_dto "github.com/root9464/Ton-students/module/service_module/dto"
	serv_repository "github.com/root9464/Ton-students/module/service_module/repository"
	"github.com/root9464/Ton-students/shared/logger"
)

type IServiceModuleService interface {
	CreateService(ctx context.Context, dto *serv_dto.ServiceType) error
	UpdateService(ctx context.Context, dto *serv_dto.UpdateServiceType) error
	GetServiceById(ctx context.Context, id string) (*serv_dto.ServiceType, error)
}

type serviceModuleService struct {
	logger    *logger.Logger
	validator *validator.Validate

	repo serv_repository.IServiceModuleRepository
}

func NewServiceModuleService(logger *logger.Logger, validator *validator.Validate, repo serv_repository.IServiceModuleRepository) *serviceModuleService {
	return &serviceModuleService{logger: logger, validator: validator, repo: repo}
}
