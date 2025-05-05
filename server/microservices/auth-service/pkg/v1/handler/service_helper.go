package handler

import (
	"errors"
	"os"
	"time"

	"github.com/QUDUSKUNLE/microservices/shared/dto"
	"github.com/golang-jwt/jwt/v5"
)

// Constants for messages
const (
	AllFields                          = "Please provide all fields"
	IncorrectPassword                  = "Incorrect passwords"
	ProvideID                          = "Id is required"
	NinRequired                        = "Nin is required"
	UserRegisteredSuccessfully         = "User registered successfully."
	OrganizationRegisteredSuccessfully = "Organization registered successfully."
	NinUpdatedSuccessfully             = "Nin updated successfully."
	WelcomeHome                        = "Welcome to Scan Records scanrecords.com."
)

// CustomContext holds the current user information
type CustomContext struct {
	User dto.CurrentUser `json:"user"`
}

// JwtCustomClaims defines the structure for JWT claims
type JwtCustomClaims struct {
	ID       string `json:"id"`
	UserType string `json:"user_type"`
	jwt.RegisteredClaims
}

// transformToken generates a JWT token for the given user
func (srv *AuthServiceStruct) transformToken(user dto.CurrentUser) (string, error) {
	secret := os.Getenv("JWT_SECRET_KEY")
	if secret == "" {
		ErrMissingSecretKey := errors.New("missing JWT secret key")
		return "", ErrMissingSecretKey
	}

	claims := &JwtCustomClaims{
		ID:       user.ID,
		UserType: user.UserType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return token, nil
}
