package user_controller

import (
	"github.com/gofiber/fiber/v2"
)

func (c *userController) Ping(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello, World!")

}
