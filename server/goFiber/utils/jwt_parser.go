package utils

import (
	"os"
	"strings"
	"github.com/golang-jwt/jwt"
	"github.com/gofiber/fiber/v2"
)

type TokenMetadata struct {
	Expires int64
}

func ExtractTokenMetadata(con *fiber.Ctx) (*TokenMetadata, error) {
	token, err := verifyToken(con)
	if err != nil {
		return nil, err
	}

	// Setting and checking toekn and credentials
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		expires := int64(claims["exp"].(float64))
		return &TokenMetadata{
			Expires: expires,
		}, nil
	}
	return nil, err
}

func extractToken(con *fiber.Ctx) string {
	bearerToken := con.Get("Authorization")

	token := strings.Split(bearerToken, " ")
	if len(token) == 2 {
		return token[1]
	}
	return ""
}

func verifyToken(con *fiber.Ctx) (*jwt.Token, error) {
	tokenString := extractToken(con)

	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("JWT_SECRET_KEY")), nil
}

