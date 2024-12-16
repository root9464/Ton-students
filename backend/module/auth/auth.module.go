package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/root9464/Ton-students/backend/module/auth/controller"
)

func AuthRoutes(router fiber.Router) {
	auth := router.Group("auth")
	auth.Get("/hello", controller.Hello)
}
