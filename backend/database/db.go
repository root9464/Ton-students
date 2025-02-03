package database

import (
	"time"

	"github.com/root9464/Ton-students/shared/logger"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

func ConnectDb(url string, log *logger.Logger) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  url,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Error),
	})

	if err != nil {
		return nil, err
	}

	log.Info("✅ Database connection successfully")

	log.Info("📦 Setting database connection pool...")
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}
