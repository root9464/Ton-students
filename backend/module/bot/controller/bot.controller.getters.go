package bot_controller

import (
	"github.com/gofiber/fiber/v2"
	bot_dto "github.com/root9464/Ton-students/module/bot/dto"
)

func (c *botController) Payment(ctx *fiber.Ctx) error {
	body := new(bot_dto.Payment)
	if err := ctx.BodyParser(body); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	paymentLink, err := c.botCommand.GeneratePayment(c.bot, body.UserId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate payment link",
		})
	}

	return ctx.JSON(fiber.Map{
		"payment_link": paymentLink,
	})
}
