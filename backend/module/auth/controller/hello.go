package controller

import "github.com/gofiber/fiber/v2"

func (c *Controller) Hello(ctx *fiber.Ctx) error {
	message := c.authService.Hello()
	return ctx.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   message,
	})
}
