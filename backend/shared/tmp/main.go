package main

import (
	"fmt"

	"github.com/root9464/Ton-students/database"
	user_model "github.com/root9464/Ton-students/module/user/model"
	"github.com/root9464/Ton-students/shared/logger"
)

const (
	USER_ID = 99281932
)

func main() {
	db, err := database.ConnectDb("postgresql://postgres.drrsenlocmidswbgjwzc:qwertyqwest7q8q1579@aws-0-eu-central-1.pooler.supabase.com:6543/postgres", logger.GetLogger())
	if err != nil {
		fmt.Println("Error connecting to database:", err)
	}

	res := db.Model(&user_model.User{}).Where("id = ?", USER_ID).Update("role", user_model.UserRole)
	if res.Error != nil {
		fmt.Println("Error updating user:", res.Error)
	}

	if res.RowsAffected == 0 {
		fmt.Println("User not found")
	}

	fmt.Println("User updated successfully")
}
