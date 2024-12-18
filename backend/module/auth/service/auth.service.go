package auth_service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/root9464/Ton-students/config"
	"github.com/root9464/Ton-students/module/auth/dto"
	"github.com/root9464/Ton-students/shared/logger"
)

var _ IAuthService = (*authService)(nil)

type IAuthService interface {
	Authorize(ctx context.Context, dto *dto.AutorizeDto) error
}

type authService struct {
	logger    *logger.Logger
	validator *validator.Validate
	config    *config.Config
}

func NewAuthService(
	logger *logger.Logger,
	validator *validator.Validate,
	config *config.Config,
) *authService {
	return &authService{
		logger:    logger,
		validator: validator,
		config:    config,
	}
}
