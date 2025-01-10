package auth_module

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/root9464/Ton-students/config"
	auth_controller "github.com/root9464/Ton-students/module/auth/controller"
	auth_service "github.com/root9464/Ton-students/module/auth/service"
	jwt_module "github.com/root9464/Ton-students/module/jwt"
	user_service "github.com/root9464/Ton-students/module/user/service"
	"github.com/root9464/Ton-students/shared/logger"
)

type AuthModule struct {
	authService    auth_service.IAuthService
	authController auth_controller.IAuthController

	logger    *logger.Logger
	validator *validator.Validate
	config    *config.Config

	userService user_service.IUserService

	jwtModule *jwt_module.JwtModule
}

func NewAuthModule(logger *logger.Logger, validator *validator.Validate, config *config.Config, userService user_service.IUserService, jwtModule *jwt_module.JwtModule) *AuthModule {
	if jwtModule == nil {
		panic("JwtModule must not be nil")
	}
	return &AuthModule{
		logger:      logger,
		validator:   validator,
		config:      config,
		userService: userService,
		jwtModule:   jwtModule,
	}
}

func (m *AuthModule) AuthService() auth_service.IAuthService {
	if m.authService == nil {
		m.authService = auth_service.NewAuthService(m.logger, m.validator, m.config, m.userService)
	}
	return m.authService
}

func (m *AuthModule) AuthController() auth_controller.IAuthController {
	if m.authController == nil {
		m.authController = auth_controller.NewAuthController(m.AuthService(), m.jwtModule)
	}
	return m.authController
}

func (m *AuthModule) AuthRoutes(router fiber.Router) {
	auth := router.Group("/auth")
	auth.Post("/authorize", m.AuthController().Authorize)
	auth.Get("/jwt-ping", m.AuthController().JwtPing)
}
