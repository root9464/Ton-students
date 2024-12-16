package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/root9464/Ton-students/backend/module/auth"
)

func AllRoutes(app *fiber.App) {
	api := app.Group("api")

	api.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Pong")
	})

	auth.AuthRoutes(api)

}
