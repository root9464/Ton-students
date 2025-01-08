package database

import (
	"time"

	"github.com/joho/godotenv"
	"google.golang.org/appengine/log"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

func ConnectDb(url string) (*gorm.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Errorf("Error loading .env file: %v", err)
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  url,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Warn),
	})

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}
