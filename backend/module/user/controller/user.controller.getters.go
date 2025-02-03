package user_controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/root9464/Ton-students/shared/utils"
)

func (c *userController) GetUser(ctx *fiber.Ctx) error {
	userInfo := ctx.Query("id")
	if userInfo == "" {
		return ctx.Status(400).JSON(&fiber.Error{
			Code:    400,
			Message: "ID is required",
		})
	}

	userIntId, err := strconv.ParseInt(userInfo, 10, 64)
	if err != nil {
		return ctx.Status(400).JSON(&fiber.Error{
			Code:    500,
			Message: "Failed convert ID",
		})
	}

	user, err := c.userService.GetUser(ctx.Context(), userIntId)
	if err != nil {
		if errorResponse, code := utils.HandlerError(err); errorResponse != nil {
			return ctx.Status(code).JSON(errorResponse)
		}
	}

	return ctx.Status(200).JSON(&fiber.Map{
		"status":  "success",
		"message": "User get successfully",
		"data":    user,
	})
}
