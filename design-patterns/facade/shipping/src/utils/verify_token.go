package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/google/uuid"
)

type ObtainedUser struct {
	ID uuid.UUID `json:"id"`
	UserType string `json:"user_type"`
}

func obtainUser(context echo.Context) *ObtainedUser {
	user := context.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	return &ObtainedUser{
		ID: claims.ID,
		UserType: claims.UserType,
	}
}

func (util *Utils) ParseUserID(context echo.Context) (*ObtainedUser, error) {
	result := obtainUser(context)
	_, err := uuid.Parse(result.ID.String())
	if err != nil {
		return &ObtainedUser{}, err
	}
	return result, nil
}
