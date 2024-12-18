package main

import (
	"github.com/gofiber/fiber/v2/log"
	_ "github.com/lib/pq"
	"github.com/root9464/Ton-students/core/app"
)

func main() {
	a := app.NewApp()
	if err := a.Run(); err != nil {
		log.Fatalf("✖ Failed to run app: %s", err.Error())
	}
}
