package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/root9464/Ton-students/backend/config"
	"github.com/root9464/Ton-students/backend/module/auth/controller"
	"github.com/root9464/Ton-students/backend/module/auth/service"
	"github.com/root9464/Ton-students/backend/shared/logger"
)

type AuthModule struct {
	authService service.IAuthService
	controller  *controller.Controller

	logger *logger.Logger
	config *config.Config
}

func NewAuthModule(config *config.Config, logger *logger.Logger) *AuthModule {
	return &AuthModule{
		logger: logger,
		config: config,
	}
}

func (m *AuthModule) AuthService() service.IAuthService {
	if m.authService == nil {
		m.authService = service.NewService(m.config, m.logger)
	}
	return m.authService
}

func (m *AuthModule) AuthController() *controller.Controller {
	if m.controller == nil {
		m.controller = controller.NewController(m.AuthService())
	}
	return m.controller
}

func (m *AuthModule) Rotes(r fiber.Router) {
	auth := r.Group("/auth")
	auth.Get("/hello", m.AuthController().Hello)
}
