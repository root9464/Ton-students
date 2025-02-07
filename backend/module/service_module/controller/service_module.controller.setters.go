package serv_controller

import (
	"github.com/gofiber/fiber/v2"
	serv_dto "github.com/root9464/Ton-students/module/service_module/dto"
	"github.com/root9464/Ton-students/shared/utils"
)

func (c *serviceModuleController) CreateService(ctx *fiber.Ctx) error {
	dto := new(serv_dto.ServiceType)

	if err := ctx.BodyParser(dto); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	if err := c.service.CreateService(ctx.Context(), dto); err != nil {
		if errorResponse, code := utils.HandlerError(err); errorResponse != nil {
			return ctx.Status(code).JSON(errorResponse)
		}
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Service created successfully",
	})
}

func (c *serviceModuleController) UpdateService(ctx *fiber.Ctx) error {
	dto := new(serv_dto.UpdateServiceType)

	if err := ctx.BodyParser(dto); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	if err := c.service.UpdateInformation(ctx.Context(), dto); err != nil {
		if errorResponse, code := utils.HandlerError(err); errorResponse != nil {
			return ctx.Status(code).JSON(errorResponse)
		}
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Service updated successfully",
	})
}
