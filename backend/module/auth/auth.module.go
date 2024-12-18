package auth_module

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/root9464/Ton-students/config"
	auth_controller "github.com/root9464/Ton-students/module/auth/controller"
	auth_service "github.com/root9464/Ton-students/module/auth/service"
	"github.com/root9464/Ton-students/shared/logger"
)

type AuthModule struct {
	authService    auth_service.IAuthService
	authController auth_controller.IAuthController

	logger    *logger.Logger
	validator *validator.Validate
	config    *config.Config
}

func NewAuthModule(
	logger *logger.Logger,
	validator *validator.Validate,
	config *config.Config,
) *AuthModule {
	return &AuthModule{
		logger:    logger,
		validator: validator,
		config:    config,
	}
}

func (m *AuthModule) AuthService() auth_service.IAuthService {
	if m.authService == nil {
		m.authService = auth_service.NewAuthService(m.logger, m.validator, m.config)
	}
	return m.authService
}

func (m *AuthModule) AuthController() auth_controller.IAuthController {
	if m.authController == nil {
		m.authController = auth_controller.NewAuthController(m.AuthService())
	}
	return m.authController
}

func (m *AuthModule) AuthRoutes(router fiber.Router) {
	auth := router.Group("/auth")
	auth.Post("/authorize", m.AuthController().Authorize)
}
