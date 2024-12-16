package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/root9464/Ton-students/backend/config"
	"github.com/root9464/Ton-students/backend/ent"
	"github.com/root9464/Ton-students/backend/module/auth"
	"github.com/root9464/Ton-students/backend/shared/utils"
)

func AllRoutes(app *fiber.App, log *utils.Logger, envs *config.Config, db *ent.Client) {
	api := app.Group("api")

	api.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Pong")
	})

	auth.AuthRoutes(api, log, envs, db)

}
