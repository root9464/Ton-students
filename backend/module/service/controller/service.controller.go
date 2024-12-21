package service_controller

import (
	"github.com/gofiber/fiber/v2"
	service_serv "github.com/root9464/Ton-students/module/service/service"
)

type IServiceController interface {
	CreateService(ctx *fiber.Ctx) error
}

type serviceController struct {
	serviceService service_serv.IServService
}

func NewServiceController(
	serviceService service_serv.IServService,
) *serviceController {
	return &serviceController{
		serviceService: serviceService,
	}
}
