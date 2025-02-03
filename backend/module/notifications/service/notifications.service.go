package notifications_service

import (
	"context"

	"github.com/go-playground/validator/v10"
	notifications_dto "github.com/root9464/Ton-students/module/notifications/dto"
	notifications_repository "github.com/root9464/Ton-students/module/notifications/repository"
	user_repository "github.com/root9464/Ton-students/module/user/repository"
	"github.com/root9464/Ton-students/shared/logger"
	"gorm.io/gorm"
)

var _ INotificationsService = (*notificationsService)(nil)

type INotificationsService interface {
	CreateNotification(ctx context.Context, notification *notifications_dto.CreateNotificationType) error

	GetNotificationsUser(ctx context.Context, userID int64) ([]notifications_dto.NotificationType, error)
}

type notificationsService struct {
	logger    *logger.Logger
	validator *validator.Validate
	db        *gorm.DB

	notificationsRepo notifications_repository.INotificationsRepository

	userRepo user_repository.IUserRepository
}

func NewNotificationsService(
	logger *logger.Logger, validator *validator.Validate, db *gorm.DB,
	notificationsRepo notifications_repository.INotificationsRepository, userRepo user_repository.IUserRepository) *notificationsService {
	return &notificationsService{logger: logger, validator: validator, db: db, notificationsRepo: notificationsRepo, userRepo: userRepo}
}
