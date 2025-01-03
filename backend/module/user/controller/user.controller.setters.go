package user_controller

import (
	"github.com/gofiber/fiber/v2"
	user_dto "github.com/root9464/Ton-students/module/user/dto"
)

func (c *userController) CreateDemo(ctx *fiber.Ctx) error {
	userData := new(user_dto.CreateUserDto)
	if err := ctx.BodyParser(userData); err != nil {
		return ctx.Status(400).JSON(&fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	user, err := c.userService.Create(ctx.Context(), userData)
	if err != nil {
		fiberErr := err.(*fiber.Error)
		return ctx.Status(fiberErr.Code).JSON(&fiber.Map{
			"status":  "failed",
			"message": fiberErr.Message,
		})
	}

	return ctx.Status(200).JSON(&fiber.Map{
		"status":  "success",
		"message": "Authorized",
		"data":    &user,
	})
}
