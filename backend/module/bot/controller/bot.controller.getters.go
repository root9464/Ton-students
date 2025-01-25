package bot_controller

import (
	"github.com/gofiber/fiber/v2"
)

func (c *botController) Payment(ctx *fiber.Ctx) error {
	userId := ctx.Query("userId")
	if userId == "" {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "User ID is required",
		})
	}

	paymentLink, err := c.botCommand.GeneratePayment(c.bot, userId)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": "Failed to generate payment link",
		})
	}

	return ctx.JSON(fiber.Map{
		"payment_link": paymentLink,
	})
}
