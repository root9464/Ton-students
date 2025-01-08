package serv_service

import (
	"github.com/go-playground/validator/v10"
	serv_repository "github.com/root9464/Ton-students/module/service_module/repository"
	"github.com/root9464/Ton-students/shared/logger"
)

type IServiceModuleService interface {
}

type serviceModuleService struct {
	logger    *logger.Logger
	validator *validator.Validate

	repo serv_repository.IServiceModuleRepository
}

func NewServiceModuleService(logger *logger.Logger, validator *validator.Validate, repo serv_repository.IServiceModuleRepository) *serviceModuleService {
	return &serviceModuleService{logger: logger, validator: validator, repo: repo}
}
