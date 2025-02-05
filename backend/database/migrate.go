package database

import (
	chat_model "github.com/root9464/Ton-students/module/chat/model"
	common_model "github.com/root9464/Ton-students/module/model/common"
	notifications_model "github.com/root9464/Ton-students/module/notifications/model"
	serv_model "github.com/root9464/Ton-students/module/service_module/model"
	user_model "github.com/root9464/Ton-students/module/user/model"
	"github.com/root9464/Ton-students/shared/logger"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB, trigger bool, log *logger.Logger) error {

	if trigger {
		log.Info("📦 Migrating database...")
		models := []interface{}{
			&user_model.User{},
			&user_model.UserInfo{},

			&serv_model.Service{},
			&serv_model.ServiceInfo{},
			&serv_model.Tags{},
			&serv_model.ServiceSettings{},

			&notifications_model.Notification{},

			&chat_model.Message{},
			&chat_model.Chat{},

			&common_model.ChatUser{},
		}

		log.Info("📦 Creating types...")

		db.Exec("CREATE TYPE selected_name AS ENUM('firstname', 'lastname', 'nickname', 'username')")
		db.Exec("CREATE TYPE role AS ENUM('administarator', 'user', 'creator', 'moderator')")
		db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
		db.Exec("CREATE TYPE notification_type AS ENUM('info', 'event', 'invite', 'comment', 'message', 'like', 'dislike')")

		if err := db.AutoMigrate(models...); err != nil {
			log.Errorf("✖ Failed to migrate database: %v", err)
			return err
		}
	}

	log.Info("✅ Database connection successfully")
	return nil
}
