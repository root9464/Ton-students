package user_controller

import (
	"github.com/gofiber/fiber/v2"
	user_service "github.com/root9464/Ton-students/module/user/service"
)

type IUserController interface {
	GetByID(ctx *fiber.Ctx) error
}

type userController struct {
	userService user_service.IUserService
}

func NewUserController(
	userService user_service.IUserService,
) *userController {
	return &userController{
		userService: userService,
	}
}
