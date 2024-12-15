package routes

import (
	ent "root/database"
	ping_routes "root/modules/ping/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func Routes(app *fiber.App, db *ent.Client, log *logrus.Logger) {
	api := app.Group("/api")

	ping_routes.Routes(api)

}
