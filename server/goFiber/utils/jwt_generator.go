package utils

import (
	"os"
	"strconv"
	"time"
	"github.com/golang-jwt/jwt"
)

func GenerateNewAccessToken() (string, error) {
	// Get secret from .env
	secret := os.Getenv("JWT_SECRET_KEY")

	// Set expires minute
	minutesCount, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))

	// Create a new claims
	claims := jwt.MapClaims{}

	// set public claims
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()

	// Create a new JWT access token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate token
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, nil 
}
