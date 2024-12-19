package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

// TelegramConfig структура для конфигурации Telegram бота
type TelegramConfig struct {
	Token string `yaml:"token"`
}

// Config общая структура для конфигурации
type Config struct {
	Telegram TelegramConfig `yaml:"telegram"`
}

// loadConfig читает конфигурационный файл и возвращает структуру Config
func LoadConfig(filePath string) (*Config, error) {
	// Открываем файл
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Декодируем YAML в структуру
	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
