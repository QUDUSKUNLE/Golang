package utils

import (
	"os"
	"time"

	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Utils struct {}

type JwtCustomClaims struct {
	ID uuid.UUID `json:"id"`
	UserType string `json:"user_type"`
	jwt.RegisteredClaims
}

func (util *Utils) GenerateAccessToken(user domain.User) (string, error) {
	// Get JWT_SECRET_KEY
	secret := os.Getenv("JWT_SECRET_KEY")
	// Create a new claims
	claims := &JwtCustomClaims{
		user.ID,
		string(user.UserType),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	// Create a new JWT access token with claims
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate token
	token, err := jwtToken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}
