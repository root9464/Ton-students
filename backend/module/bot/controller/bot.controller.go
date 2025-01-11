package bot_controller

// import (
// 	"github.com/gofiber/fiber"

// 	service "github.com/root9464/Ton-students/module/bot/service"


// )
// func (c *BotController) GeneratePaymentHandler(ctx *fiber.Ctx) error {

// 	// Вызываем метод GeneratePayment
// 	paymentLink, err := service.GeneratePayment(bot, ctx, c.logger.Logger)
// 	if err != nil {
// 		c.logger.Error("Failed to generate payment link: " + err.Error())
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": err.Error(),
// 		})
// 	}

// 	// Возвращаем ссылку на оплату
// 	return ctx.JSON(fiber.Map{
// 		"payment_link": paymentLink,
// 	})
// }
