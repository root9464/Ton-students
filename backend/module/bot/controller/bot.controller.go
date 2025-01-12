package bot_controller

import (
	"log"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/gofiber/fiber/v2"

	dto "github.com/root9464/Ton-students/module/bot/dto"
	service "github.com/root9464/Ton-students/module/bot/service"
)

func (c *BotController) GeneratePaymentHandler(bot *gotgbot.Bot, ctx *fiber.Ctx) error {
	body := new(dto.Payment)

	if err := ctx.BodyParser(body); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Failed to parse request body",
        })
    }
	log.Printf("user id: %s\n", body.UserId)

	paymentLink, err := service.GeneratePayment(bot, body.UserId, c.logger.Logger)
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
