package models

import (
	"time"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)

const (
	USER    UserType = "USER"
	CARRIER UserType = "CARRIER"
	ADMIN   UserType = "ADMIN"
)

type (
	User struct {
		gorm.Model
		ID        uuid.UUID       `json:"id" gorm:"uuid;primaryKey"`
		CreatedAt *time.Time      `json:"created_at"`
		UpdatedAt *time.Time      `json:"updated_at"`
		DeletedAt *gorm.DeletedAt `gorm:"index" json:"-"`
		Email     string          `json:"email" gorm:"unique"`
		Nin       string          `json:"nin" gorm:"unique"`
		Password  string          `json:"password"`
		UserType  UserType        `json:"user_type"`
	}
	UserType string
	UserDto  struct {
		Email           string   `json:"email" validate:"email,required"`
		Password        string   `json:"password" validate:"min=8,required"`
		ConfirmPassword string   `json:"confirm_password" validate:"eqfield=Password"`
		UserType        UserType `json:"user_type" validate:"oneof=USER CARRIER,required"`
	}
	LogInDto struct {
		Email    string `json:"email" validate:"email,required"`
		Password string `json:"password" validate:"min=8,required"`
	}
	ResetPasswordDto struct {
		Email string `json:"email" validate:"email,required"`
	}
	CurrentUser struct {
		ID uuid.UUID `json:"id"`
		UserType string `json:"user_type"`
	}
	Response struct {
		Data    interface{} `json:"data"`
		Success bool        `json:"success"`
	}
	ErrorResponse struct {
		Error   interface{} `json:"error"`
		Success bool        `json:"success"`
	}
)

func (user UserType) ReturnUserString() string {
	switch user {
	case USER:
		return string(USER)
	case CARRIER:
		return string(CARRIER)
	}
	return "Unknown"
}

func (user *User) BeforeCreate(scope *gorm.DB) error {
	user.ID = uuid.New()
	pass, _ := hashPassword(user.Password)
	user.Password = pass
	user.UserType = UserType(user.UserType)
	return nil
}

func (u *User) BuildNewUser(user UserDto) (*User, error) {
	Password, err := hashPassword(user.Password)
	if err != nil {
		return &User{}, err
	}
	return &User{
		Email:    user.Email,
		Password: Password,
		UserType: UserType(user.UserType),
	}, nil
}

func (user *User) ComparePassword(pass string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass)); err != nil {
		return errors.New("incorrect log in credentials")
	}
	return nil
}

func hashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	password = string(hashPassword)
	return password, nil
}
