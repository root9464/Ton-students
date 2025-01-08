package serv_controller

import (
	"github.com/gofiber/fiber/v2"
	serv_service "github.com/root9464/Ton-students/module/service_module/service"
)

type IServiceModuleController interface {
	Pong(ctx *fiber.Ctx) error
}

type serviceModuleController struct {
	service serv_service.IServiceModuleService
}

func NewServiceModuleController(service serv_service.IServiceModuleService) *serviceModuleController {
	return &serviceModuleController{service: service}
}
