package user_module

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/root9464/Ton-students/database"
	user_controller "github.com/root9464/Ton-students/module/user/controller"
	user_repository "github.com/root9464/Ton-students/module/user/repository"
	user_service "github.com/root9464/Ton-students/module/user/service"
	"github.com/root9464/Ton-students/shared/logger"
)

type UserModule struct {
	userService    user_service.IUserService
	userController user_controller.IUserController
	userRepo       user_repository.IUserRepository

	logger    *logger.Logger
	validator *validator.Validate
	db        *database.Database
}

func NewUserModule(
	logger *logger.Logger,
	validator *validator.Validate,
	db *database.Database,
) *UserModule {
	return &UserModule{
		logger:    logger,
		validator: validator,
		db:        db,
	}
}

func (m *UserModule) UserRepo() user_repository.IUserRepository {
	if m.userRepo == nil {
		m.userRepo = user_repository.NewUserRepository(m.db, m.logger)
	}
	return m.userRepo
}

func (m *UserModule) UserService() user_service.IUserService {
	if m.userService == nil {
		m.userService = user_service.NewUserService(m.logger, m.validator, m.UserRepo())
	}
	return m.userService
}

func (m *UserModule) UserController() user_controller.IUserController {
	if m.userController == nil {
		m.userController = user_controller.NewUserController(m.UserService())
	}
	return m.userController
}

func (m *UserModule) UserRoutes(router fiber.Router) {
	user := router.Group("/user")

	user.Post("/create", m.UserController().CreateDemo)
}
