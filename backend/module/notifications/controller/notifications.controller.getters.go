package notifications_controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/root9464/Ton-students/shared/utils"
)

func (c *notificationsController) GetNotificationsUser(ctx *fiber.Ctx) error {
	userId := ctx.Query("id")

	if userId == "" {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "User ID is required",
		})
	}

	userIntId, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return ctx.Status(400).JSON(&fiber.Error{
			Code:    500,
			Message: "Failed convert ID",
		})
	}

	notifications, err := c.notificationsService.GetNotificationsUser(ctx.Context(), userIntId)
	if err != nil {
		if errorResponse, code := utils.HandlerError(err); errorResponse != nil {
			return ctx.Status(code).JSON(errorResponse)
		}
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Notifications get successfully",
		"data":    notifications,
	})
}
