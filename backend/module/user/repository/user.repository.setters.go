package user_repository

import (
	"context"
	"fmt"

	user_model "github.com/root9464/Ton-students/module/user/model"
)

func (r *userRepository) Create(ctx context.Context, user *user_model.User) (*user_model.User, error) {
	r.logger.Info("Creating user...")
	if err := r.db.WithContext(ctx).Create(&user).Preload("Infos").Error; err != nil {
		r.logger.Errorf("Error creating user: %v", err)
		return nil, err
	}
	r.logger.Info("User create successfully")
	return user, nil
}

func (r *userRepository) Update(ctx context.Context, user *user_model.User) (*user_model.User, error) {
	r.logger.Info("Updating user...")
	result := r.db.WithContext(ctx).Model(&user_model.User{}).Where("id = ?", user.ID).Updates(user)
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
	if err := r.db.WithContext(ctx).Create(&userInfo).Error; err != nil {
		r.logger.Errorf("Error adding user info: %v", err)
		return err
	}
	r.logger.Info("User info added successfully")
	return nil
}

func (r *userRepository) UpdateUserInfo(ctx context.Context, userInfo *user_model.UserInfo) error {
	r.logger.Info("Updating user info...")

	result := r.db.WithContext(ctx).Model(&user_model.UserInfo{}).Where("id = ?", userInfo.ID).Updates(userInfo)
	if result.Error != nil {
		r.logger.Errorf("Error updating user info: %v", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		err := fmt.Errorf("user info with ID '%s' does not exist", userInfo.ID)
		r.logger.Warnf("%v", err)
		return err
	}

	r.logger.Info("User info updated successfully")
	return nil
}

func (r *userRepository) DeleteUserInfo(ctx context.Context, userInfoID string) error {
	r.logger.Info("Deleting user info...")

	result := r.db.WithContext(ctx).Where("id = ?", userInfoID).Delete(&user_model.UserInfo{})

	if result.Error != nil {
		r.logger.Errorf("Error deleting user info: %v", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		r.logger.Errorf("Error deleting user info: user info with ID %s not found", userInfoID)
		return fmt.Errorf("user info with ID %s not found", userInfoID)
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

func (r *userRepository) ChangeUserRole(ctx context.Context, id string, role string) error {
	r.logger.Info("Changing user role...")

	result := r.db.WithContext(ctx).Model(&user_model.User{}).Where("id = ?", id).Update(role, user_model.UserRole)
	if result.Error != nil {
		r.logger.Errorf("Error changing user role: %v", result.Error)
		return fmt.Errorf("error changing user role")
	}

	if result.RowsAffected == 0 {
		r.logger.Errorf("Error changing user role: user with ID %s not found", id)
		return fmt.Errorf("user with ID %s not found", id)
	}

	r.logger.Infof("Successfully changed user role for user with ID %s", id)
	return nil
}
