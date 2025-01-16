package bot_controller

import (
	"log"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/gofiber/fiber/v2"

	dto "github.com/root9464/Ton-students/module/bot/dto"
)

func (c *BotController) GeneratePaymentHandler(bot *gotgbot.Bot, ctx *fiber.Ctx) error {
	body := new(dto.Payment)

	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}
	log.Printf("user id: %s\n", body.UserId)

	paymentLink, err := c.service.GeneratePayment(bot, body.UserId)
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

func (c *BotController) InvateLinkHandler(bot *gotgbot.Bot, ctx *fiber.Ctx) error {
	invateLink, err := c.service.InvateLink(bot)
	if err != nil {
		c.logger.Error("Failed to generate invate link: " + err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate invate link",
		})
	}

	return ctx.JSON(fiber.Map{
		"invate_link": invateLink,
	})
}
