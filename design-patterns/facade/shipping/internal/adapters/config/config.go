package config

import (
	"github.com/QUDUSKUNLE/shipping/internal/core/utils"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/joho/godotenv"
)

func LoadEnvironmentVariable() error {
	if err := godotenv.Load(".env"); err != nil {
		return err
	}
	return nil
}

func JWTConfig(secret string) echojwt.Config {
	return echojwt.Config{
		NewClaimsFunc: func(context echo.Context) jwt.Claims {
			return new(utils.JwtCustomClaims)
		},
		SigningKey: []byte(secret),
	}
}
