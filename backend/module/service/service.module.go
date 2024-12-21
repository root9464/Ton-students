package service_module

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/root9464/Ton-students/ent"
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
	db        *ent.Client
}

func NewServiceModule(
	logger *logger.Logger,
	validator *validator.Validate,
	db *ent.Client,
) *ServiceModule {
	return &ServiceModule{
		logger:    logger,
		validator: validator,
		db:        db,
	}
}

func (m *ServiceModule) ServiceRepo() service_repository.IServiceRepository {
	if m.serviceRepository == nil {
		m.serviceRepository = service_repository.NewServiceRepository(m.db, m.logger)
	}
	return m.serviceRepository
}

func (m *ServiceModule) ServService() service_serv.IServService {
	if m.servService == nil {
		m.servService = service_serv.NewServService(m.logger, m.validator, m.db, m.ServiceRepo())
	}
	return m.servService
}

func (m *ServiceModule) ServiceController() service_controller.IServiceController {
	if m.serviceController == nil {
		m.serviceController = service_controller.NewServiceController(m.ServService())
	}
	return m.serviceController
}

func (m *ServiceModule) ServiceRoutes(router fiber.Router) {
	service := router.Group("/service")
	service.Post("/create", m.ServiceController().CreateService)
}
