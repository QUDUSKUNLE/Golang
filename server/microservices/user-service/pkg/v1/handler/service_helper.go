package handler

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/QUDUSKUNLE/microservices/shared/constants"
	"github.com/QUDUSKUNLE/microservices/shared/db"
	"github.com/QUDUSKUNLE/microservices/shared/dto"
	userProtoc "github.com/QUDUSKUNLE/microservices/shared/protogen/user"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Constants for messages
const (
	AllFields                          = "Please provide all fields"
	NotFound                           = "User not found"
	IncorrectPassword                  = "Incorrect passwords"
	ProvideID                          = "Id is required"
	NinRequired                        = "Nin is required"
	UserRegisteredSuccessfully         = "User registered successfully."
	OrganizationRegisteredSuccessfully = "Organization registered successfully."
	NinUpdatedSuccessfully             = "Nin updated successfully."
	WelcomeHome                        = "Welcome to Scan Records scanrecords.com."
	ErrUnauthorized                    = "Unauthorized to perform operation."
	ErrInvalidCredentials              = "Incorrect login credentials."
	ErrNinUpdated                      = "Nin updated successfully."
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

// transformUserRPC converts a CreateUserRequest to a UserDto
func (srv *UserServiceStruct) transformUserRPC(req *userProtoc.CreateUserRequest) dto.UserDto {
	return dto.UserDto{
		Password:        req.GetPassword(),
		Email:           req.GetEmail(),
		ConfirmPassword: req.GetConfirmPassword(),
		UserType:        db.UserEnum(req.GetUserType().Enum().String()),
	}
}

// transformToken generates a JWT token for the given user
func (srv *UserServiceStruct) transformToken(user dto.CurrentUser) (string, error) {
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

func transformUserToProto(user db.User) *userProtoc.User {
	return &userProtoc.User{
		Id:        user.ID,
		Email:     user.Email.String,
		CreatedAt: user.CreatedAt.Time.String(),
		UpdatedAt: user.UpdatedAt.Time.String(),
	}
}

func getUserFromContext(ctx context.Context) (*constants.UserType, error) {
	user, ok := ctx.Value("user").(*constants.UserType)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, ErrUnauthorized)
	}
	return user, nil
}
