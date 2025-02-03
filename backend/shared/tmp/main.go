package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/root9464/Ton-students/database"
	user_model "github.com/root9464/Ton-students/module/user/model"
	"github.com/root9464/Ton-students/shared/logger"
	"gorm.io/gorm"
)

const (
	USER_ID = 99281932
)

func main() {
	db, _ := database.ConnectDb("postgresql://postgres.drrsenlocmidswbgjwzc:qwertyqwest7q8q1579@aws-0-eu-central-1.pooler.supabase.com:6543/postgres", logger.GetLogger())

	user := new([]user_model.User)

	if err := db.WithContext(context.Background()).Offset(0).Limit(2).
		Preload("Services.Infos").
		Preload("Services.Tags").
		Preload("Services.Settings").
		Find(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Printf("User with ID %d not found", 1)
		}
		fmt.Printf("Error getting user by ID: %v", err)

	}

	jsonData, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		fmt.Println("Ошибка при преобразовании в JSON:", err)
		return
	}

	fmt.Println("Результат в формате JSON:")
	fmt.Println(string(jsonData))
}
