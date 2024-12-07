package handler

import (
	module_hello "root/modules/hello"

	"github.com/gofiber/fiber/v2"
)

func RoutesHandler(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/ping", module_hello.Hello)

}
