package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	TelegramBotToken string `mapstructure:"TELEGRAM_BOT_TOKEN"`
	DatabaseUrl      string `mapstructure:"DATABASE_URL"`
	KafkaBrokerUrl   string `mapstructure:"KAFKA_BROKER_URL"`
	BotToken         string `mapstructure:"BOT_APITOKEN"`
	AdminId          int64  `mapstructure:"ADMIN_ID"`
	ChannelId        int64  `mapstructure:"REQUIRED_CHANNEL"`
}

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("../.env")
	viper.SetConfigType("env")

	// Automatically map environment variables
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}
	return nil
}
