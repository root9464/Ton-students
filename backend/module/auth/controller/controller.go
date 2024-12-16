package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/root9464/Ton-students/backend/module/auth/service"
)

func Hello(ctx *fiber.Ctx) error {

	message := service.SayHello()

	return ctx.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   message,
	})
}
