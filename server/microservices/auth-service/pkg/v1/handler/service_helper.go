package handler

import (
	"os"
	"time"

	"github.com/QUDUSKUNLE/microservices/auth-service/adapters/db"
	"github.com/QUDUSKUNLE/microservices/auth-service/adapters/dto"
	userProtoc "github.com/QUDUSKUNLE/microservices/auth-service/protogen/golang/user"
	"github.com/golang-jwt/jwt/v5"
)

const (
	All_Fields              = "Please provide all fields"
	Incorrect_Password      = "Incorrect passwords"
	Provide_ID              = "Id is required"
	Not_Found               = "User`s not found"
	Nin_Required            = "Nin is required"
	Registered_Successfully = "User registered successfully."
	Welcome_Home            = "Welcome to Bahsoon Shipping Inc."
)

type CustomContext struct {
	User dto.CurrentUser `json:"user"`
}

type JwtCustomClaims struct {
	// Authorized bool       `json:"authorized"`
	ID       string          `json:"id"`
	UserType db.NullUserEnum `json:"user_type"`
	jwt.RegisteredClaims
}

func (srv *UserServiceStruct) transformUserRPC(req *userProtoc.CreateUserRequest) dto.UserDto {
	return dto.UserDto{Password: req.GetPassword(), Email: req.GetEmail(), ConfirmPassword: req.GetConfirmPassword(), UserType: db.UserEnum(req.GetUserType().String())}
}

func (srv *UserServiceStruct) transformUsers(us []*db.User) *userProtoc.GetUsersResponse {
	users := make([]*userProtoc.User, 0)
	for _, user := range us {
		users = append(users, &userProtoc.User{
			Email:     user.Email.String,
			CreatedAt: user.CreatedAt.Time.String(),
			UpdatedAt: user.UpdatedAt.Time.String(),
		})
	}
	return &userProtoc.GetUsersResponse{Data: users}
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
