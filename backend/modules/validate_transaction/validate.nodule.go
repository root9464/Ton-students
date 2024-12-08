package validate_transaction

import "github.com/gofiber/fiber/v2"

func Validate(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"message": "hello world",
	})
}
