package notifications_service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	notifications_dto "github.com/root9464/Ton-students/module/notifications/dto"
	notifications_model "github.com/root9464/Ton-students/module/notifications/model"
	"github.com/samber/lo"
)

func (s *notificationsService) GetNotificationsUser(ctx context.Context, userID int64) ([]notifications_dto.NotificationType, error) {
	s.logger.Infof("Getting notifications for user: %d", userID)
	if err := s.validator.Var(userID, "required"); err != nil {
		s.logger.Warnf("validate error: %s", err.Error())
		return nil, &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	userNotifications, err := s.notificationsRepo.GetNotificationsUser(ctx, userID)
	if err != nil {
		s.logger.Errorf("Error getting notifications for user: %v", err)
		return nil, err
	}

	notificationsDTO := lo.Map(userNotifications, func(notification notifications_model.Notification, _ int) notifications_dto.NotificationType {
		sender, err := s.userRepo.GetByID(ctx, notification.SenderID)
		if err != nil {
			s.logger.Errorf("Error getting sender for notification: %v", err)
			return notifications_dto.NotificationType{}
		}
		receiver, err := s.userRepo.GetByID(ctx, notification.ReceiverID)
		if err != nil {
			s.logger.Errorf("Error getting receiver for notification: %v", err)
			return notifications_dto.NotificationType{}
		}

		return notifications_dto.NotificationType{
			Sender:   sender.Username,
			Receiver: receiver.Username,
			Content:  notification.Content,
			Type:     string(notification.Type),
			Created:  notification.CreatedAt.Format("2006-01-02 15:04:05"),
		}
	})

	if len(notificationsDTO) == 0 {
		s.logger.Info("empty notifications")
		return nil, nil
	}

	return notificationsDTO, nil
}
