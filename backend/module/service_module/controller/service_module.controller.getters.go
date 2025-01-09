package serv_controller

import (
	"github.com/gofiber/fiber/v2"
)

func (c *serviceModuleController) Pong(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(&fiber.Map{"msg": "pong "})
}

func (c *serviceModuleController) GetAllServices(ctx *fiber.Ctx) error {
	services, err := c.repo.GetAllServices(ctx.Context())
	if err != nil {
		return ctx.Status(500).JSON(&fiber.Error{
			Code:    500,
			Message: err.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":   "success",
		"message":  "Services fetched successfully",
		"services": services,
	})
}

func (c *serviceModuleController) GetServiceById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if id == "" {
		return ctx.Status(400).JSON(&fiber.Error{
			Code:    400,
			Message: "ID is required",
		})
	}

	service, err := c.service.GetServiceById(ctx.Context(), id)
	if err != nil {
		return ctx.Status(500).JSON(&fiber.Error{
			Code:    500,
			Message: err.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Service fetched successfully",
		"service": service,
	})
}
