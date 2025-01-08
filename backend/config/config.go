package config

import (
	"fmt"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	TelegramBotToken string `mapstructure:"TELEGRAM_BOT_TOKEN"`
	DatabaseUrl      string `mapstructure:"DATABASE_URL"`
	RedisUrl         string `mapstructure:"REDIS_URL"`
}

func validateConfig(config *Config) error {
	configMap := map[string]string{
		"TELEGRAM_BOT_TOKEN": config.TelegramBotToken,
		"DATABASE_URL":       config.DatabaseUrl,
		"REDIS_URL":          config.RedisUrl,
	}

	for key, value := range configMap {
		if strings.TrimSpace(value) == "" {
			return fmt.Errorf("missing required configuration field: %s", key)
		}
	}

	return nil
}

func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	err = validateConfig(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}
