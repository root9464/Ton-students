package auth_controller

import (
	"github.com/gofiber/fiber/v2"
	auth_service "github.com/root9464/Ton-students/module/auth/service"
)

type IAuthController interface {
	Authorize(ctx *fiber.Ctx) error
}

type AuthController struct {
	authService auth_service.IAuthService
}

func NewAuthController(
	authService auth_service.IAuthService,
) *AuthController {
	return &AuthController{
		authService: authService,
	}
}
