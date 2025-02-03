package notifications_controller

import (
	"github.com/gofiber/fiber/v2"
	notifications_dto "github.com/root9464/Ton-students/module/notifications/dto"
	"github.com/root9464/Ton-students/shared/utils"
)

func (c *notificationsController) CreateNotification(ctx *fiber.Ctx) error {
	dto := new(notifications_dto.CreateNotificationType)

	if err := ctx.BodyParser(dto); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	if err := c.notificationsService.CreateNotification(ctx.Context(), dto); err != nil {
		if errorResponse, code := utils.HandlerError(err); errorResponse != nil {
			return ctx.Status(code).JSON(errorResponse)
		}
	}

	return ctx.Status(200).JSON(&fiber.Map{
		"status":  "success",
		"message": "Notification created successfully",
	})
}
