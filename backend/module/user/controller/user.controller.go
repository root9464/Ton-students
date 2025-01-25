package user_controller

import (
	"github.com/gofiber/fiber/v2"
	user_service "github.com/root9464/Ton-students/module/user/service"
)

type IUserController interface {
	GetUser(ctx *fiber.Ctx) error
	AddUserInfo(ctx *fiber.Ctx) error
	SelectVisibleName(ctx *fiber.Ctx) error
	SetUserNickname(ctx *fiber.Ctx) error
	UpdateUserInfo(ctx *fiber.Ctx) error
	DeleteUserInfo(ctx *fiber.Ctx) error
	AddManyUserInfo(ctx *fiber.Ctx) error
}

type userController struct {
	userService user_service.IUserService
}

func NewUserController(userService user_service.IUserService) *userController {
	return &userController{userService: userService}
}
