package user_repository

import (
	"context"

	user_model "github.com/root9464/Ton-students/module/user/model"
)

func (r *userRepository) Create(ctx context.Context, user *user_model.User) (*user_model.User, error) {
	r.logger.Info("Creating user...")
	if err := r.db.Db.Create(&user).Error; err != nil {
		r.logger.Errorf("Error creating user: %v", err)
		return nil, err
	}
	r.logger.Info("User create successfully")
	return user, nil
}

func (r *userRepository) Update(ctx context.Context, user *user_model.User) (*user_model.User, error) {
	r.logger.Info("Updating user...")
	result := r.db.Db.Model(&user_model.User{}).Where("id = ?", user.ID).Updates(user)
	if err := result.Error; err != nil {
		r.logger.Errorf("Error updating user: %v", err)
		return nil, err
	}

	if result.RowsAffected == 0 {
		return nil, ErrUserNotFound
	}

	r.logger.Info("User update successfully")
	return user, nil
}

func (r *userRepository) AddUserInfo(ctx context.Context, userInfo *user_model.UserInfo) error {
	r.logger.Info("Adding user info...")
	if err := r.db.Db.Create(&userInfo).Error; err != nil {
		r.logger.Errorf("Error adding user info: %v", err)
		return err
	}
	r.logger.Info("User info added successfully")
	return nil
}
