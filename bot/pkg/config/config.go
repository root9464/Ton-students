package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	BotToken  string
	AdminId   int64
	ChannelId int64
}

func New() *Config {
	err := godotenv.Load("./../../.env")
	if err != nil {
		panic("Error loading .env file: " + err.Error())
	}

	adminID, err := strconv.ParseInt(os.Getenv("ADMIN_ID"), 10, 64)
	if err != nil {
		panic("Error parsing ADMIN_ID: " + err.Error())
	}

	channelID, err := strconv.ParseInt(os.Getenv("REQUIRED_CHANNEL"), 10, 64)
	if err != nil {
		panic("Error parsing REQUIRED_CHANNEL: " + err.Error())
	}

	return &Config{
		BotToken:  os.Getenv("BOT_APITOKEN"),
		AdminId:   adminID,
		ChannelId: channelID,
	}
}
