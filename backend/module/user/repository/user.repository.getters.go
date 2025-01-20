package user_repository

import (
	"context"

	user_model "github.com/root9464/Ton-students/module/user/model"
	"gorm.io/gorm"
)

func (r *userRepository) GetByID(ctx context.Context, id int64) (*user_model.User, error) {
	r.logger.Info("Getting user by ID...")
	user := new(user_model.User)
	if err := r.db.WithContext(ctx).Preload("Infos").First(&user, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			r.logger.Infof("User with ID %d not found", id)
			return nil, nil
		}
		r.logger.Errorf("Error getting user by ID: %v", err)
		return nil, err
	}

	if len(user.Infos) == 0 {
		user.Infos = nil
	}

	r.logger.Infof("User with ID %d retrieved successfully", id)
	return user, nil
}

func (r *userRepository) GetByHash(ctx context.Context, hash string) (*user_model.User, error) {
	r.logger.Info("Getting user by hash...")
	user := new(user_model.User)
	if err := r.db.WithContext(ctx).Preload("Infos").First(&user, "hash = ?", hash).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			r.logger.Infof("User with hash %s not found", hash)
			return nil, nil
		}
		r.logger.Errorf("Error getting user by hash: %v", err)
		return nil, err
	}

	r.logger.Infof("User with hash %s retrieved successfully", hash)
	return user, nil
}

func (r *userRepository) UserServices(ctx context.Context) (*user_model.User, error) {
	r.logger.Info("Getting creator services...")

	user := new(user_model.User)

	if err := r.db.WithContext(ctx).Preload("Services.Infos").Preload("Services.Tags").Preload("Services.Settings").Find(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			r.logger.Infof("User with ID %d not found", 1)
			return nil, nil
		}
		r.logger.Errorf("Error getting user by ID: %v", err)
		return nil, err
	}

	return user, nil
}
