package utils

import (
	"os"
	// "strconv"
	"time"
	"github.com/google/uuid"
	"github.com/golang-jwt/jwt/v5"
	// "github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
	// echojwt "github.com/labstack/echo-jwt/v4"
)

type Utils struct {}

type JwtCustomClaims struct {
	ID uuid.UUID `json:"id"`
	jwt.RegisteredClaims
}

func (util *Utils) GenerateAccessToken(id uuid.UUID) (string, error) {
	// Get JWT_SECRET_KEY
	secret := os.Getenv("JWT_SECRET_KEY")
	// Create a new claims
	claims := &JwtCustomClaims{
		id,
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
