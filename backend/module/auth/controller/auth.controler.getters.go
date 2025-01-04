package auth_controller

import (
	"github.com/gofiber/fiber/v2"
	auth_dto "github.com/root9464/Ton-students/module/auth/dto"
)

func (c *authController) Authorize(ctx *fiber.Ctx) error {
	data := new(auth_dto.AutorizeDto)
	if err := ctx.BodyParser(data); err != nil {
		return ctx.Status(400).JSON(&fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	user, err := c.authService.Authorize(ctx.Context(), data)
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
