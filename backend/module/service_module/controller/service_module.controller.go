package serv_controller

import (
	"github.com/gofiber/fiber/v2"
	serv_repository "github.com/root9464/Ton-students/module/service_module/repository"
	serv_service "github.com/root9464/Ton-students/module/service_module/service"
)

var _ IServiceModuleController = (*serviceModuleController)(nil)

type IServiceModuleController interface {
	Pong(ctx *fiber.Ctx) error
	CreateService(ctx *fiber.Ctx) error
	UpdateService(ctx *fiber.Ctx) error

	ServiceFeed(ctx *fiber.Ctx) error
	GetServiceById(ctx *fiber.Ctx) error
}

type serviceModuleController struct {
	service serv_service.IServiceModuleService
	repo    serv_repository.IServiceModuleRepository
}

func NewServiceModuleController(service serv_service.IServiceModuleService, repo serv_repository.IServiceModuleRepository) *serviceModuleController {
	return &serviceModuleController{service: service, repo: repo}
}
