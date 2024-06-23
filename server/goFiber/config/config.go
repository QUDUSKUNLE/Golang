package config

import (
	"os"
	"fmt"
	"github.com/joho/godotenv"
)

func Config(key string) string {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}
