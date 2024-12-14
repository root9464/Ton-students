package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	BotToken string
	AdminId  int64
}

func New() *Config {
	err := godotenv.Load("./../../.env")
	if err != nil {
		panic(err.Error())
	}

	adminID, err := strconv.ParseInt(os.Getenv("ADMIN_ID"), 10, 64)
	if err != nil {
		panic(err.Error())
	}

	return &Config{
		BotToken: os.Getenv("BOT_APITOKEN"),
		AdminId:  adminID,
	}
}
