package auth_controller

import (
	"github.com/gofiber/fiber/v2"
	auth_service "github.com/root9464/Ton-students/module/auth/service"
	jwt_module "github.com/root9464/Ton-students/module/jwt"
	jwt_funcs "github.com/root9464/Ton-students/module/jwt/functions"
)

type IAuthController interface {
	Authorize(ctx *fiber.Ctx) error
	JwtPing(ctx *fiber.Ctx) error
}

type authController struct {
	authService auth_service.IAuthService
	jwtModule   jwt_funcs.IJwtFuncs
}

func NewAuthController(authService auth_service.IAuthService, jwtModule jwt_module.JwtModule) *authController {
	return &authController{
		authService: authService,
		jwtModule:   jwtModule.JwtFuncs(),
	}
}
