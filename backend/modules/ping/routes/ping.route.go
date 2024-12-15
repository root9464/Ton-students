package ping_routes

import (
	"root/modules/ping"

	"github.com/gofiber/fiber/v2"
)

func Routes(router fiber.Router) {
	router.Get("/ping", ping.PingHandler)
}
