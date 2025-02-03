package notifications_module

import (
	"github.com/go-playground/validator/v10"
	notifications_repository "github.com/root9464/Ton-students/module/notifications/repository"
	notifications_service "github.com/root9464/Ton-students/module/notifications/service"
	user_module "github.com/root9464/Ton-students/module/user"
	"github.com/root9464/Ton-students/shared/logger"
	"gorm.io/gorm"
)

type NotificationsModule struct {
	logger    *logger.Logger
	validator *validator.Validate
	db        *gorm.DB

	// notificationsController notifications_controller.INotificationsController
	notificationsService notifications_service.INotificationsService
	notificationsRepo    notifications_repository.INotificationsRepository

	userModule user_module.UserModule
}

func (m *NotificationsModule) NotificationsService() notifications_service.INotificationsService {
	if m.notificationsService == nil {
		m.notificationsService = notifications_service.NewNotificationsService(m.logger, m.validator, m.db, m.NotificationsRepository(), m.userModule.UserRepo())
	}
	return m.notificationsService
}

// func (m *NotificationsModule) NotificationsController() notifications_controller.INotificationsController {
// 	if m.notificationsController == nil {
// 		m.notificationsController = notifications_controller.NewNotificationsController(m.NotificationsService())
// 	}
// 	return m.notificationsController
// }

func (m *NotificationsModule) NotificationsRepository() notifications_repository.INotificationsRepository {
	if m.notificationsRepo == nil {
		m.notificationsRepo = notifications_repository.NewNotificationsRepository(m.db, m.logger)
	}
	return m.notificationsRepo
}

func NewNotificationsModule(
	logger *logger.Logger, validator *validator.Validate, db *gorm.DB,
	userModule user_module.UserModule,
) *NotificationsModule {
	return &NotificationsModule{
		logger: logger, validator: validator, db: db, userModule: userModule,
	}
}

// func (m *NotificationsModule) NotificationsRoutes(router fiber.Router) {
// 	notifications := router.Group("/notifications")
// 	notifications.Post("/create", m.NotificationsController().CreateNotification)
// 	notifications.Get("/get-notifications", m.NotificationsController().GetNotificationsUser)
// }
