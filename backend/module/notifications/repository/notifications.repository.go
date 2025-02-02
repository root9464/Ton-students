package notifications_repository

import (
	"context"

	notifications_model "github.com/root9464/Ton-students/module/notifications/model"
	"github.com/root9464/Ton-students/shared/logger"
	"gorm.io/gorm"
)

var _ INotificationsRepository = (*notificationsRepository)(nil)

type INotificationsRepository interface {
	CreateNotification(ctx context.Context, notification *notifications_model.Notification) error
	GetNotificationsUser(ctx context.Context, userID int64) ([]notifications_model.Notification, error)
}

type notificationsRepository struct {
	db     *gorm.DB
	logger *logger.Logger
}

func NewNotificationsRepository(db *gorm.DB, logger *logger.Logger) *notificationsRepository {
	return &notificationsRepository{
		db:     db,
		logger: logger,
	}
}
