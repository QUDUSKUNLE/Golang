package handlers

import (
	"errors"
	"os"
	"time"

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

func (httpHandler *HTTPHandler) currentUser(context echo.Context) *CurrentUser {
	user := context.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	return &CurrentUser{
		ID: claims.ID,
		UserType: claims.UserType,
	}
}

func (httpHandler *HTTPHandler) ValidateStruct(context echo.Context, obj interface{}) error {
	// Bind userDto
	if err := context.Bind(obj); err != nil {
		return err
	}
	// Validate user input
	if err := context.Validate(obj); err != nil {
		return err
	}
	return nil
}

func (httpHandler *HTTPHandler) parseUserID(context echo.Context) (*CurrentUser, error) {
	result := httpHandler.currentUser(context)
	_, err := uuid.Parse(result.ID.String())
	if err != nil {
		return &CurrentUser{}, err
	}
	return result, nil
}

func (httpHandler *HTTPHandler) PrivateMiddlewareContext(context echo.Context, userType string) (*CurrentUser, error) {
	user, err := httpHandler.parseUserID(context)
	if err != nil {
		return &CurrentUser{}, err
	}
	// Check user type
	if user.UserType != userType {
		return &CurrentUser{}, errors.New("unauthorized to perform this operation")
	}
	return user, nil
}


func (util *HTTPHandler) GenerateAccessToken(user CurrentUser) (string, error) {
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

func (util *HTTPHandler) ComputeErrorResponse(status int, message interface{}, context echo.Context) error {
	return context.JSON(status, echo.Map{
		"error": message,
		"success": false,
	})
}

func (util *HTTPHandler) ComputeResponseMessage(status int, message interface{}, context echo.Context) error {
	return context.JSON(status, echo.Map{
		"result": message,
		"success": true,
	})
}
