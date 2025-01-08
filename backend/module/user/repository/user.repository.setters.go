package user_repository

import (
	"context"
	"fmt"

	user_model "github.com/root9464/Ton-students/module/user/model"
)

func (r *userRepository) Create(ctx context.Context, user *user_model.User) (*user_model.User, error) {
	r.logger.Info("Creating user...")
	if err := r.db.Create(&user).Error; err != nil {
		r.logger.Errorf("Error creating user: %v", err)
		return nil, err
	}
	r.logger.Info("User create successfully")
	return user, nil
}

func (r *userRepository) Update(ctx context.Context, user *user_model.User) (*user_model.User, error) {
	r.logger.Info("Updating user...")
	result := r.db.Model(&user_model.User{}).Where("id = ?", user.ID).Updates(user)
	if err := result.Error; err != nil {
		r.logger.Errorf("Error updating user: %v", err)
		return nil, err
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("user not found")
	}

	r.logger.Info("User update successfully")
	return user, nil
}

func (r *userRepository) AddUserInfo(ctx context.Context, userInfo *user_model.UserInfo) error {
	r.logger.Info("Adding user info...")
	if err := r.db.Create(&userInfo).Error; err != nil {
		r.logger.Errorf("Error adding user info: %v", err)
		return err
	}
	r.logger.Info("User info added successfully")
	return nil
}

func (r *userRepository) UpdateUserInfo(ctx context.Context, userInfo *user_model.UserInfo) error {
	r.logger.Info("Updating user info...")
	if err := r.db.Model(&user_model.UserInfo{}).Where("id = ?", userInfo.ID).Updates(userInfo).Error; err != nil {
		r.logger.Errorf("Error updating user info: %v", err)
		return err
	}
	r.logger.Info("User info updated successfully")
	return nil
}

func (r *userRepository) DeleteUserInfo(ctx context.Context, userInfoID string) error {
	r.logger.Info("Deleting user info...")
	if err := r.db.Where("id = ?", userInfoID).Delete(&user_model.UserInfo{}).Error; err != nil {
		r.logger.Errorf("Error deleting user info: %v", err)
		return err
	}
	r.logger.Info("User info deleted successfully")
	return nil
}

func (r *userRepository) AddManyUserInfo(ctx context.Context, userInfo []*user_model.UserInfo) error {
	if len(userInfo) == 0 {
		return fmt.Errorf("no user info to add")
	}

	r.logger.Info("Adding user info...")

	result := r.db.WithContext(ctx).CreateInBatches(userInfo, 100)
	if result.Error != nil {
		r.logger.Errorf("Error adding user info: %v", result.Error)
		return fmt.Errorf("error adding user info")
	}

	r.logger.Infof("Successfully added %d user info records", result.RowsAffected)
	return nil
}
