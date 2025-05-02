package handler

import (
	"encoding/json"
	"fmt"

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
		Password:        req.GetPassword(),
		Email:           req.GetEmail(),
		ConfirmPassword: req.GetConfirmPassword(),
		UserType:        db.UserEnum(req.GetUserType().Enum().String()),
	}
}

func transformUserToProto(user db.User) *userProtoc.User {
	// Initialize default address and contact
	address := &dto.Address{}
	contact := &dto.Contact{}

	// Handle non-existing or invalid address
	if len(user.Address) > 0 {
		if err := json.Unmarshal(user.Address, address); err != nil {
			fmt.Printf("Failed to unmarshal address: %v\n", err)
			address = &dto.Address{} // Use an empty address if unmarshaling fails
		}
	}

	// Handle non-existing or invalid contact
	if len(user.Contact) > 0 {
		if err := json.Unmarshal(user.Contact, contact); err != nil {
			fmt.Printf("Failed to unmarshal contact: %v\n", err)
			contact = &dto.Contact{} // Use an empty contact if unmarshaling fails
		}
	}

	// Transform user to proto format
	return &userProtoc.User{
		Id:    user.ID,
		Email: user.Email.String,
		Address: &userProtoc.Address{
			Street:  address.Street,
			City:    address.City,
			State:   address.State,
			Country: address.Country,
		},
		Contact: &userProtoc.Contact{
			Phone: func(phones []string) []*userProtoc.Phone {
				var protoPhones []*userProtoc.Phone
				for _, phone := range phones {
					protoPhones = append(protoPhones, &userProtoc.Phone{Phone: phone})
				}
				return protoPhones
			}(contact.Phone),
			Email: contact.Email,
		},
		Nin:       user.Nin.String,
		CreatedAt: user.CreatedAt.Time.String(),
		UpdatedAt: user.UpdatedAt.Time.String(),
	}
}
