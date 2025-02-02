package notifications_controller

import (
	"github.com/gofiber/fiber/v2"
	notifications_service "github.com/root9464/Ton-students/module/notifications/service"
)

var _ INotificationsController = (*notificationsController)(nil)

type INotificationsController interface {
	CreateNotification(ctx *fiber.Ctx) error
	GetNotificationsUser(ctx *fiber.Ctx) error
}

type notificationsController struct {
	notificationsService notifications_service.INotificationsService
}

func NewNotificationsController(notificationsService notifications_service.INotificationsService) *notificationsController {
	return &notificationsController{notificationsService: notificationsService}
}
