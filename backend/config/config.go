package config

import "github.com/gofiber/fiber/v2/middleware/cors"

var CORS_CONFIG = cors.New(
	cors.Config{
		AllowOrigins:     "http://localhost:5173, http://0.0.0.0:5173, https://4f67-95-105-125-55.ngrok-free.app",
		AllowCredentials: true,
	},
)
