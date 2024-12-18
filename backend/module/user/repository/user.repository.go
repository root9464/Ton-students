package user_repository

import (
	"context"

	"github.com/root9464/Ton-students/ent"
	"github.com/root9464/Ton-students/shared/logger"
	"github.com/telegram-mini-apps/init-data-golang"
)

var _ IUserRepository = (*userRepository)(nil)

type IUserRepository interface {
	Create(ctx context.Context, user *initdata.InitData) error
	GetByID(ctx context.Context, id int64) (*ent.User, error)
	Update(ctx context.Context, user *initdata.InitData) error
}

type userRepository struct {
	db     *ent.Client
	logger *logger.Logger
}

func NewUserRepository(
	db *ent.Client,
	logger *logger.Logger,
) *userRepository {
	return &userRepository{
		db:     db,
		logger: logger,
	}
}
