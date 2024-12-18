package auth_controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/root9464/Ton-students/module/auth/dto"
)

func (c *AuthController) Authorize(ctx *fiber.Ctx) error {
	user := new(dto.AutorizeDto)

	if err := ctx.BodyParser(user); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Failed to parse body",
		}
	}

	if err := c.authService.Authorize(ctx.Context(), user); err != nil {
		return ctx.JSON(&fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	return ctx.JSON(&fiber.Map{
		"status": "success",
	})
}
