package user_controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (c *userController) GetByID(ctx *fiber.Ctx) error {
	id, _ := strconv.ParseInt(ctx.Params("id"), 10, 64)

	user, err := c.userService.GetByID(ctx.Context(), id)
	if err != nil {
		fiberErr := err.(*fiber.Error)
		return ctx.Status(fiberErr.Code).JSON(&fiber.Map{
			"status":  "failed",
			"message": fiberErr.Message,
		})
	}

	return ctx.Status(200).JSON(&fiber.Map{
		"status": "success",
		"data":   user,
	})
}
