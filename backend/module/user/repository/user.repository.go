package user_repository

import (
	"context"
	"errors"

	user_model "github.com/root9464/Ton-students/module/user/model"
	"github.com/root9464/Ton-students/shared/database"
	"github.com/root9464/Ton-students/shared/logger"
)

var _ IUserRepository = (*userRepository)(nil)
var ErrUserNotFound = errors.New("user not found")

type IUserRepository interface {
	Create(ctx context.Context, user *user_model.User) error
	GetByID(ctx context.Context, id int64) (*user_model.User, error)
	Update(ctx context.Context, user *user_model.User) error
}

type userRepository struct {
	db     *database.Database
	logger *logger.Logger
}

func NewUserRepository(
	db *database.Database,
	logger *logger.Logger,
) *userRepository {
	return &userRepository{
		db:     db,
		logger: logger,
	}
}
