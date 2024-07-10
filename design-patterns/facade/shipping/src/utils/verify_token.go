package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/google/uuid"
)


func obtainUser(context echo.Context) string {
	user := context.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	ID := claims.ID
	return ID.String()
}

func (util *Utils) ParseUserID(context echo.Context) (uuid.UUID, error) {
	return uuid.Parse(obtainUser(context))
}
