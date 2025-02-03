package notifications_repository

import (
	"context"

	notifications_model "github.com/root9464/Ton-students/module/notifications/model"
)

func (r *notificationsRepository) CreateNotification(ctx context.Context, notification *notifications_model.Notification) error {
	r.logger.Info("Creating notification...")

	if err := r.db.WithContext(ctx).Create(&notification).Error; err != nil {
		r.logger.Errorf("Error creating notification: %v", err)
		return err
	}

	return nil
}
