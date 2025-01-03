package auth_service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/root9464/Ton-students/config"
	auth_dto "github.com/root9464/Ton-students/module/auth/dto"
	user_model "github.com/root9464/Ton-students/module/user/model"
	user_service "github.com/root9464/Ton-students/module/user/service"
	"github.com/root9464/Ton-students/shared/logger"
)

var _ IAuthService = (*authService)(nil)

type IAuthService interface {
	Authorize(ctx context.Context, dto *auth_dto.AutorizeDto) (*user_model.User, error)
}

type authService struct {
	logger      *logger.Logger
	validator   *validator.Validate
	config      *config.Config
	userService user_service.IUserService
}

func NewAuthService(
	logger *logger.Logger,
	validator *validator.Validate,
	config *config.Config,
	userService user_service.IUserService,
) *authService {
	return &authService{
		logger:      logger,
		validator:   validator,
		config:      config,
		userService: userService,
	}
}
