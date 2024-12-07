package main

import (
	"log"
	"root/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()
	app.Use(cors.New(
		cors.Config{
			AllowOrigins:     "http://localhost:5173, http://0.0.0.0:5173, https://4f67-95-105-125-55.ngrok-free.app",
			AllowCredentials: true,
		},
	))

	handler.RoutesHandler(app)

	log.Fatal(app.Listen(":3000"))
}
