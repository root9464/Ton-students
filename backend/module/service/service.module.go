package service_module

import (
	"github.com/go-playground/validator/v10"
	"github.com/root9464/Ton-students/config"
	service_controller "github.com/root9464/Ton-students/module/service/controller"
	service_repository "github.com/root9464/Ton-students/module/service/repository"
	service_serv "github.com/root9464/Ton-students/module/service/service"
	"github.com/root9464/Ton-students/shared/logger"
)

type ServiceModule struct {
	servService       service_serv.IServService
	serviceController service_controller.IServiceController
	serviceRepository service_repository.IServiceRepository

	logger    *logger.Logger
	validator *validator.Validate
	config    *config.Config
}

func NewServiceModule(
	logger *logger.Logger,
	validator *validator.Validate,
	config *config.Config,
) *ServiceModule {
	return &ServiceModule{
		logger:    logger,
		validator: validator,
		config:    config,
	}
}
