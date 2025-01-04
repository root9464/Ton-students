package database

import (
	"github.com/gofiber/fiber/v2/log"
	serv_model "github.com/root9464/Ton-students/module/service_module/model"
	user_model "github.com/root9464/Ton-students/module/user/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB, trigger bool) error {

	if trigger {
		log.Info("📦 Migrating database...")
		models := []interface{}{
			&user_model.User{},
			&user_model.UserInfo{},
			&serv_model.Service{},
			&serv_model.ServiceInfo{},
			&serv_model.Tags{},
		}

		db.Exec("CREATE TYPE selected_name AS ENUM('firstname', 'lastname', 'nickname', 'username')")
		db.Exec("CREATE TYPE role AS ENUM('administarator', 'user', 'creator', 'moderator')")
		db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
		if err := db.AutoMigrate(models...); err != nil {
			return err
		}
	}

	log.Info("✅ Database connection successfully")
	return nil
}
