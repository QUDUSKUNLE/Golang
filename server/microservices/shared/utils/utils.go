package utils

import (
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	KafkaBroker string
	KafkaTopic string
	KafkaGroup string
	DB_URL string
}

func LoadEnvironmentVariable() error {
	if err := godotenv.Load(".env"); err != nil {
		return err
	}
	return nil
}

func LoadConfig() (*Config, error) {
	if err := LoadEnvironmentVariable(); err != nil {
		return nil, err
	}

	config := &Config{
		Port:        os.Getenv("PORT"),
		KafkaBroker: os.Getenv("KAFKA_BROKER"),
		KafkaTopic:  os.Getenv("KAFKA_TOPIC"),
		KafkaGroup:  os.Getenv("KAFKA_GROUP_ID"),
		DB_URL: os.Getenv("DB_URL"),
	}
	return config, nil
}
