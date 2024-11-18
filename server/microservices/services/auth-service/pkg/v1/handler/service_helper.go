package handler

import (
	"os"
	"time"
	"github.com/google/uuid"
	"github.com/golang-jwt/jwt/v5"
	"github.com/QUDUSKUNLE/microservices/services/auth-service/internal/models"
	userProtoc "github.com/QUDUSKUNLE/microservices/services/auth-service/protogen/golang/user"
)

type CustomContext struct {
	User models.CurrentUser `json:"user"`
}

type JwtCustomClaims struct {
	ID uuid.UUID `json:"id"`
	UserType string `json:"user_type"`
	jwt.RegisteredClaims
}

func (srv *UserServiceStruct) transformUserRPC(req *userProtoc.CreateUserRequest) models.User {
	return models.User{Password: req.GetPassword(), Email: req.GetEmail(), UserType: models.UserType(req.GetUserType())}
}

func (srv *UserServiceStruct) transformUsers(us []*models.User) *userProtoc.GetUsersResponse {
	users := make([]*userProtoc.User, 0)
	for _, user := range us {
		users = append(users, &userProtoc.User{
			Id:    user.ID.String(),
			Email: user.Email,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
		})
	}
	return &userProtoc.GetUsersResponse{Data: users}
}

func (srv *UserServiceStruct) transformToken(user models.CurrentUser) (string, error) {
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
