package serv_controller

import "github.com/gofiber/fiber/v2"

func (c *serviceModuleController) Pong(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(&fiber.Map{"msg": "pong"})
}
