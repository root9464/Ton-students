package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/root9464/Ton-students/backend/config"
	"github.com/root9464/Ton-students/backend/ent"
	auth_controller "github.com/root9464/Ton-students/backend/module/auth/controller"
	"github.com/root9464/Ton-students/backend/shared/utils"
)

func AuthRoutes(router fiber.Router, log *utils.Logger, envs *config.Config, db *ent.Client) {
	auth := router.Group("auth")

	auth.Post("/authorize", auth_controller.Authorize(log, envs, db))
}
