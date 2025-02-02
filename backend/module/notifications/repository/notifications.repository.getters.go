package notifications_repository

import (
	"context"

	notifications_model "github.com/root9464/Ton-students/module/notifications/model"
	"gorm.io/gorm"
)

func (r *notificationsRepository) GetNotificationsUser(ctx context.Context, userID int64) ([]notifications_model.Notification, error) {
	r.logger.Info("Getting notifications for user...")

	notifications := new([]notifications_model.Notification)
	if err := r.db.WithContext(ctx).Where("receiver_id = ?", userID).Find(&notifications).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			r.logger.Info("No notifications found")
			return nil, nil
		}
		r.logger.Errorf("Error getting notifications: %v", err)
		return nil, err
	}

	return *notifications, nil
}
