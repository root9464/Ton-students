package main

import (
	_ "github.com/lib/pq"

	"github.com/gofiber/fiber/v2"
	config "github.com/root9464/Ton-students/backend/config"
	"github.com/root9464/Ton-students/backend/core/routes"
	"github.com/root9464/Ton-students/backend/ent"
	"github.com/root9464/Ton-students/backend/shared/middleware"
	"github.com/root9464/Ton-students/backend/shared/utils"
)

func main() {
	const port = ":6069"

	log := utils.GetLogger()
	globalConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	app.Use(config.CORS_CONFIG)
	app.Use(middleware.LoggerMiddleware())

	_, err = ent.Open("postgres", globalConfig.DatabaseUrl)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	log.Info("connected to the database successfully")
	routes.AllRoutes(app)

	log.Infof("🌐 Server is running on %s", port)
	log.Info("✅ Server started successfully")
	if err := app.Listen(port); err != nil {
		log.Fatal("❌ Server startup error:", err)
	}
}
