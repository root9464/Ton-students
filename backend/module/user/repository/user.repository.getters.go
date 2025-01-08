package user_repository

import (
	"context"
	"time"

	user_model "github.com/root9464/Ton-students/module/user/model"
	"gorm.io/gorm"
)

func (r *userRepository) GetByID(ctx context.Context, id int64) (*user_model.User, error) {
	r.logger.Info("Getting user...")

	user := new(user_model.User)

	if err := r.db.Db.Preload("Infos").First(&user, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		r.logger.Errorf("Error getting user: %v", err)
		return nil, err
	}

	r.logger.Info("User get successfully")
	return user, nil
}

func (r *userRepository) GetByHash(ctx context.Context, hash string) (*user_model.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond) // Тайм-аут 100ms
	defer cancel()

	r.logger.Info("Getting user...")

	user := new(user_model.User)

	err := r.db.Db.WithContext(ctx).
		Select("id, role, hash"). // Загружаем только нужные поля
		First(&user, "hash = ?", hash).
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		r.logger.Errorf("Error getting user: %v", err)
		return nil, err
	}

	r.logger.Info("User retrieved successfully")
	return user, nil
}
