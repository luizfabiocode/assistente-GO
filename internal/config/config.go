package config

import (
	"errors"
	"os"
)

type Config struct {
	OpenRouterAPIKey string
	ModelName        string
	TelegramToken    string
}

func Load() (*Config, error) {
	apiKey := os.Getenv("OPENROUTER_API_KEY")
	if apiKey == "" {
		return nil, errors.New("a variável de ambiente OPENROUTER_API_KEY não está definida")
	}

	telegramToken := os.Getenv("TELEGRAM_TOKEN")
	if telegramToken == "" {
		return nil, errors.New("a variável de ambiente TELEGRAM_TOKEN não está definida")
	}

	modelName := os.Getenv("MODEL_NAME")

	return &Config{
		OpenRouterAPIKey: apiKey,
		ModelName:        modelName,
		TelegramToken:    telegramToken,
	}, nil
}
