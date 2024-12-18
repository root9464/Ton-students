package auth_controller

import (
	"github.com/gofiber/fiber/v2"
	auth_service "github.com/root9464/Ton-students/module/auth/service"
)

type IAuthController interface {
	Authorize(ctx *fiber.Ctx) error
}

type authController struct {
	authService auth_service.IAuthService
}

func NewAuthController(
	authService auth_service.IAuthService,
) *authController {
	return &authController{
		authService: authService,
	}
}
