package service_controller

import (
	"github.com/gofiber/fiber/v2"
	service_dto "github.com/root9464/Ton-students/module/service/dto"
)

func (c *serviceController) CreateService(ctx *fiber.Ctx) error {
	data := new(service_dto.CreateServiceDto)
	if err := ctx.BodyParser(data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "failed",
			"message": "Invalid request body: " + err.Error(),
		})
	}

	service, err := c.serviceService.CreateService(ctx.Context(), data)
	if err != nil {
		if fiberErr, ok := err.(*fiber.Error); ok {
			return ctx.Status(fiberErr.Code).JSON(&fiber.Map{
				"status":  "failed",
				"message": fiberErr.Message,
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "failed",
			"message": "Internal server error",
		})
	}

	// Возвращаем успешный ответ со статусом 201
	return ctx.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status": "success",
		"data":   service,
	})
}
