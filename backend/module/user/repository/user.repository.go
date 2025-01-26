package user_repository

import (
	"context"

	user_model "github.com/root9464/Ton-students/module/user/model"
	"github.com/root9464/Ton-students/shared/logger"
	"gorm.io/gorm"
)

var _ IUserRepository = (*userRepository)(nil)

type IUserRepository interface {
	GetByID(ctx context.Context, id int64) (*user_model.User, error)
	GetByHash(ctx context.Context, hash string) (*user_model.User, error)

	Create(ctx context.Context, user *user_model.User) (*user_model.User, error)
	Update(ctx context.Context, user *user_model.User) (*user_model.User, error)
	AddUserInfo(ctx context.Context, userInfo *user_model.UserInfo) error
	UpdateUserInfo(ctx context.Context, userInfo *user_model.UserInfo) error
	DeleteUserInfo(ctx context.Context, userInfoID string) error
	AddManyUserInfo(ctx context.Context, userInfo []*user_model.UserInfo) error

	UserServices(ctx context.Context) (*[]user_model.User, error)
}

type userRepository struct {
	db     *gorm.DB
	logger *logger.Logger
}

func NewUserRepository(
	db *gorm.DB,
	logger *logger.Logger,
) *userRepository {
	return &userRepository{
		db:     db,
		logger: logger,
	}
}
