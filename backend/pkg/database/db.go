package database

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type Database struct {
	Db *gorm.DB
}

var DB Database

func ConnectDb(url string) (Database, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Warn),
	})
	if err != nil {
		return Database{}, err
	}

	DB = Database{Db: db}

	sqlDB, err := DB.Db.DB()
	if err != nil {
		return Database{}, err
	}

	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return DB, nil
}
