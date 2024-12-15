package main

import (
	"root/config"
	"root/core/routes"
	"root/shared/lib"
	"root/shared/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	const port = ":6069"

	log := utils.Logger()

	_, err := lib.ConnectDatabase(log)
	if err != nil {
		log.WithError(err).Fatal("Failed to connect to the database")
	}

	app := fiber.New()
	app.Use(config.CORS_CONFIG)

	db, err := lib.ConnectDatabase(log)
	if err != nil {
		log.WithError(err).Fatal("Failed to connect to the database")
	}

	routes.Routes(app, db, log)

	if err := app.Listen(port); err != nil {
		log.WithError(err).Fatal("Failed to start the server")
	}

}
