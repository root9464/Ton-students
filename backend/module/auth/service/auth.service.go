package auth_service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/root9464/Ton-students/config"
	auth_dto "github.com/root9464/Ton-students/module/auth/dto"
	jwt_funcs "github.com/root9464/Ton-students/module/jwt/functions"
	user_dto "github.com/root9464/Ton-students/module/user/dto"
	user_service "github.com/root9464/Ton-students/module/user/service"
	"github.com/root9464/Ton-students/shared/logger"
)

var _ IAuthService = (*authService)(nil)

type IAuthService interface {
	Authorize(ctx context.Context, dto *auth_dto.AutorizeDto) (*user_dto.ShortUserType, error)
}

type authService struct {
	logger    *logger.Logger
	validator *validator.Validate
	config    *config.Config

	userService user_service.IUserService
	jwtFuncs    jwt_funcs.IJwtFuncs
}

func NewAuthService(
	logger *logger.Logger,
	validator *validator.Validate,
	config *config.Config,
	userService user_service.IUserService,
	jwtFuncs jwt_funcs.IJwtFuncs,
) *authService {
	return &authService{
		logger:      logger,
		validator:   validator,
		config:      config,
		userService: userService,
		jwtFuncs:    jwtFuncs,
	}
}
