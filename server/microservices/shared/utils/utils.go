package utils

import (
	"github.com/joho/godotenv"
)

func LoadEnvironmentVariable() error {
	if err := godotenv.Load(".env"); err != nil {
		return err
	}
	return nil
}
