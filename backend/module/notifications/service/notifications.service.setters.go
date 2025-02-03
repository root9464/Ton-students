package notifications_service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	notifications_dto "github.com/root9464/Ton-students/module/notifications/dto"
	notifications_model "github.com/root9464/Ton-students/module/notifications/model"
	"github.com/root9464/Ton-students/shared/utils"
)

func (s *notificationsService) CreateNotification(ctx context.Context, notification *notifications_dto.CreateNotificationType) error {

	s.logger.Infof("dto received: %+v", notification)
	if err := s.validator.Struct(notification); err != nil {
		s.logger.Warnf("validate error: %s", err.Error())
		return &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}
	s.logger.Infof("Validating dto success : %+v", notification)

	s.logger.Infof("converting dto to entity: %+v", notification)
	newNotification, err := utils.ConvertDtoToEntity[notifications_model.Notification](notification)
	if err != nil {
		s.logger.Warnf("convert dto to entity error: %s", err.Error())
		return &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	if err := s.notificationsRepo.CreateNotification(ctx, newNotification); err != nil {
		s.logger.Warnf("create notification error: %s", err.Error())
		return &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}
	return nil
}
