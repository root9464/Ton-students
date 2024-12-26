package user_service

import (
	"context"

	"github.com/go-playground/validator/v10"
	user_model "github.com/root9464/Ton-students/module/user/model"
	user_repository "github.com/root9464/Ton-students/module/user/repository"
	"github.com/root9464/Ton-students/shared/logger"
)

var _ IUserService = (*userService)(nil)

type IUserService interface {
	Create(ctx context.Context, dto interface{}) error
	GetByID(ctx context.Context, id int64) (*user_model.User, error)
}

type userService struct {
	logger    *logger.Logger
	validator *validator.Validate

	repo user_repository.IUserRepository
}

func NewUserService(
	logger *logger.Logger,
	validator *validator.Validate,
	repo user_repository.IUserRepository,
) *userService {
	return &userService{
		logger:    logger,
		validator: validator,
		repo:      repo,
	}
}
