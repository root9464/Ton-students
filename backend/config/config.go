package config

import (
	"fmt"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseUrl string `mapstructure:"DATABASE_URL"`
	RedisUrl    string `mapstructure:"REDIS_URL"`

	JwtPrivateKey string `mapstructure:"JWT_PRIVATE_KEY"`
	JwtPublicKey  string `mapstructure:"JWT_PUBLIC_KEY"`

	TelegramBotToken        string `mapstructure:"TELEGRAM_BOT_TOKEN"`
	TelegramAdminId         int64  `mapstructure:"TELEGRAM_ADMIN_ID"`
	TelegramRequiredChannel int64  `mapstructure:"TELEGRAM_REQUIRED_CHANNEL"`
	TelegramChatBotId       int64  `mapstructure:"TELEGRAM_CHAT_BOT_ID"`
}

func validateConfig(config *Config) error {
	configMap := map[string]interface{}{
		"DATABASE_URL": config.DatabaseUrl,
		"REDIS_URL":    config.RedisUrl,

		"JWT_PRIVATE_KEY": config.JwtPrivateKey,
		"JWT_PUBLIC_KEY":  config.JwtPublicKey,

		"TELEGRAM_BOT_TOKEN":        config.TelegramBotToken,
		"TELEGRAM_ADMIN_ID":         config.TelegramAdminId,
		"TELEGRAM_REQUIRED_CHANNEL": config.TelegramRequiredChannel,
		"TELEGRAM_CHAT_BOT_ID":      config.TelegramChatBotId,
	}

	for key, value := range configMap {
		if isEmptyValue(value) {
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

func isEmptyValue(value interface{}) bool {
	switch v := value.(type) {
	case string:
		return strings.TrimSpace(v) == ""
	case int64:
		return v == 0
	case nil:
		return true
	default:
		return false
	}
}
