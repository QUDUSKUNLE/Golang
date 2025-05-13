package handler

import (
	"github.com/QUDUSKUNLE/microservices/shared/db"
	"github.com/QUDUSKUNLE/microservices/shared/dto"
	userProtoc "github.com/QUDUSKUNLE/microservices/shared/protogen/user"
	"github.com/golang-jwt/jwt/v5"
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
	DiagnosticRegisteredSuccessfully   = "Diagnostic registered successfully."
	UserUpdatedSuccessfully            = "User information updated successfully."
	WelcomeHome                        = "Welcome to S3records."
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
		Email:            req.GetEmail(),
		Password:         req.GetPassword(),
		ConfirmPassword:  req.GetConfirmPassword(),
		DiagnosticCentre: req.GetDiagnosticCentre(),
		UserType:         db.UserEnum(req.GetUserType().Enum().String()),
	}
}

func transformUserToProto(user db.User) *userProtoc.User {
	// Transform user to proto format
	return &userProtoc.User{
		Id:        user.ID,
		Email:     user.Email.String,
		Nin:       user.Nin.String,
		CreatedAt: user.CreatedAt.Time.String(),
		UpdatedAt: user.UpdatedAt.Time.String(),
	}
}
