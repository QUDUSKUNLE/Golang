package handlers

import (
		"github.com/golang-jwt/jwt/v5"
		"github.com/google/uuid"
		"github.com/labstack/echo/v4"
)

type CurrentUser struct {
	ID uuid.UUID `json:"id"`
	UserType string `json:"user_type"`
}

type JwtCustomClaims struct {
	ID uuid.UUID `json:"id"`
	UserType string `json:"user_type"`
	jwt.RegisteredClaims
}

func (httpHandler *HTTPHandler) CurrentUser(context echo.Context) *CurrentUser {
	user := context.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	return &CurrentUser{
		ID: claims.ID,
		UserType: claims.UserType,
	}
}

func (httpHandler *HTTPHandler) ParseUserID(context echo.Context) (*CurrentUser, error) {
	result := httpHandler.CurrentUser(context)
	_, err := uuid.Parse(result.ID.String())
	if err != nil {
		return &CurrentUser{}, err
	}
	return result, nil
}
