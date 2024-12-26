package database

import (
	"github.com/gofiber/fiber/v2/log"
	user_model "github.com/root9464/Ton-students/module/user/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	log.Info("📦 Migrating database...")
	models := []interface{}{
		&user_model.User{},
	}

	db.Exec("CREATE TYPE selected_name AS ENUM('firstname', 'lastname', 'nickname', 'username')")
	db.Exec("CREATE TYPE role AS ENUM('administarator', 'user', 'creator', 'moderator')")

	if err := db.AutoMigrate(models...); err != nil {
		return err
	}
	return nil
}
