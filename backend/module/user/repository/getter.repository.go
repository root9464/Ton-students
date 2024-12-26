package user_repository

import (
	"context"

	user_model "github.com/root9464/Ton-students/module/user/model"
	"gorm.io/gorm"
)

func (r *userRepository) GetByID(ctx context.Context, id int64) (*user_model.User, error) {
	r.logger.Info("Getting user...")
	user := new(user_model.User)
	if err := r.db.Db.Where("id = ?", id).Find(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		r.logger.Errorf("Error getting user: %v", err)
		return nil, err
	}
	r.logger.Info("User get successfully")
	return user, nil
}
