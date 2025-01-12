package bot_controller

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/gofiber/fiber/v2"

	service "github.com/root9464/Ton-students/module/bot/service"
)

func (c *BotController) GeneratePaymentHandler(bot *gotgbot.Bot, ctx *fiber.Ctx) error {
	paymentLink, err := service.GeneratePayment(bot, c.logger.Logger)
	if err != nil {
		c.logger.Error("Failed to generate payment link: " + err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate payment link",
		})
	}

	return ctx.JSON(fiber.Map{
		"payment_link": paymentLink,
	})
}
