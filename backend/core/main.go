package main

import (
	"github.com/root9464/Ton-students/backend/config"
	"github.com/root9464/Ton-students/backend/core/app"
	"github.com/root9464/Ton-students/backend/shared/lib"
	"github.com/root9464/Ton-students/backend/shared/logger"
)

func main() {
	// const port = ":6069"
	//
	log := logger.GetLogger()
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("✖ Failed to load config: %s", err.Error())
	}
	log.Info("✔  Config loaded")

	_, err = lib.ConnectDatabase(log.Logger)
	//
	if err != nil {
		log.WithError(err).Fatal("Failed to connect to the database")
	}

	_, err = lib.ConnectDatabase(log.Logger)
	if err != nil {
		log.WithError(err).Fatal("Failed to connect to the database")
	}

	app, err := app.NewApp(&config, log)
	if err != nil {
		log.WithError(err).Fatal("Failed to create the app")
	}

	if err := app.Run(); err != nil {
		log.WithError(err).Fatal("Failed to run the app")
	}
}
