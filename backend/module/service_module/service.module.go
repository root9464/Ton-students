package service_module

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	jwt_module "github.com/root9464/Ton-students/module/jwt"
	serv_controller "github.com/root9464/Ton-students/module/service_module/controller"
	serv_repository "github.com/root9464/Ton-students/module/service_module/repository"
	serv_service "github.com/root9464/Ton-students/module/service_module/service"
	user_repository "github.com/root9464/Ton-students/module/user/repository"
	"github.com/root9464/Ton-students/shared/logger"
	"github.com/root9464/Ton-students/shared/middleware"
	"gorm.io/gorm"
)

type ServiceModule struct {
	service           serv_service.IServiceModuleService
	serviceController serv_controller.IServiceModuleController
	serviceRepo       serv_repository.IServiceModuleRepository

	logger    *logger.Logger
	validator *validator.Validate
	db        *gorm.DB
	jwtModule jwt_module.JwtModule
	publicKey string
	userRepo  user_repository.IUserRepository
}

func NewServiceModule(
	logger *logger.Logger, validator *validator.Validate, db *gorm.DB,
	userRepo user_repository.IUserRepository, jwtModule jwt_module.JwtModule, publicKey string,
) *ServiceModule {
	return &ServiceModule{logger: logger, validator: validator, db: db, userRepo: userRepo, jwtModule: jwtModule, publicKey: publicKey}
}

func (m *ServiceModule) ServiceRepo() serv_repository.IServiceModuleRepository {
	if m.serviceRepo == nil {
		m.serviceRepo = serv_repository.NewServiceModuleRepository(m.db, m.logger, m.userRepo)
	}
	return m.serviceRepo
}

func (m *ServiceModule) ServService() serv_service.IServiceModuleService {
	if m.service == nil {
		m.service = serv_service.NewServiceModuleService(m.logger, m.validator, m.ServiceRepo(), m.db)
	}
	return m.service
}

func (m *ServiceModule) ServiceController() serv_controller.IServiceModuleController {
	if m.serviceController == nil {
		m.serviceController = serv_controller.NewServiceModuleController(m.ServService(), m.ServiceRepo())
	}
	return m.serviceController
}

func (m *ServiceModule) ServiceRoutes(router fiber.Router) {
	middleware := middleware.NewRoleMiddleware(m.logger, m.userRepo, m.jwtModule.JwtHelpers(), m.publicKey)

	service := router.Group("/service", middleware.CreatorOnly())

	service.Get("/ping", m.ServiceController().Pong)
	service.Post("/create", m.ServiceController().CreateService)
	service.Get("/get-service", m.ServiceController().GetServiceById)

	service.Get("/feed", m.ServiceController().ServiceFeed)

	service.Patch("/update", m.ServiceController().UpdateInformation)
}
