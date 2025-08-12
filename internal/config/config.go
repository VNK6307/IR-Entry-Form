package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	BotToken string
}

func LoadConfig() (*Config, error) {

	_ = godotenv.Load()

	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		return nil, errors.New("BOT_TOKEN environment variable not set")
	}

	return &Config{
			token,
		},
		nil
}
