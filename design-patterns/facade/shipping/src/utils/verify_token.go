package utils

import (
	"github.com/labstack/echo/v4"
	"github.com/golang-jwt/jwt/v5"
)


func (util *Utils) ObtainUser(context echo.Context) string {
	user := context.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	ID := claims.ID
	return ID.String()
}