package handler

import (
	"os"
	"time"
	"github.com/QUDUSKUNLE/microservices/hospital-service/adapters/dto"
	// userProtoc "github.com/QUDUSKUNLE/microservices/hospital-service/protogen/golang/user"
	"github.com/golang-jwt/jwt/v5"
)

const (
	All_Fields                           = "Please provide all fields"
	Incorrect_Password                   = "Incorrect passwords"
	Provide_ID                           = "Id is required"
	Not_Found                            = "User`s not found"
	Nin_Required                         = "Nin is required"
	User_Registered_Successfully         = "User registered successfully."
	Organization_Registered_Successfully = "Organization registered successfully."
	Welcome_Home                         = "Welcome to Scan Records scanrecords.com."
)

type CustomContext struct {
	User dto.CurrentUser `json:"user"`
}

type JwtCustomClaims struct {
	// Authorized bool       `json:"authorized"`
	ID       string `json:"id"`
	UserType string `json:"user_type"`
	jwt.RegisteredClaims
}


func (srv *UserServiceStruct) transformToken(user dto.CurrentUser) (string, error) {
	// Get JWT_SECRET_KEY
	secret := os.Getenv("JWT_SECRET_KEY")
	// Create a new claims
	claims := &JwtCustomClaims{
		// Authorized: true,
		user.ID,
		user.UserType,
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
