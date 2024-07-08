package model

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserType string

const (
	USER UserType = "USER"
	RIDER UserType = "RIDER"
	UNKNOWN UserType = "UNKNOWN"
)

type User struct {
	gorm.Model
	ID 		   		uuid.UUID `db:"id"`
	Email 	 		string    `db:"email"`
	Password		string    `db:"pass"`
	CreatedAt 	time.Time `db:"created_at"`
	UpdatedAt 	time.Time `db:"updated_at"`
	UserType    UserType  `db:"user_type"`
}

type UserDTO struct {
	Email    				string `json:"email" binding:"required,email,lte=100" validate:"required"`
	Pass 				    string `json:"password" binding:"required,gte=6,lte=20" validate:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required,gte=6,lte=20" validate:"required"`
	UserType 				string 	`json:"userType" binding:"required" validate:"required"`
}

type LogInDTO struct {
	Email  string `json:"email" binding:"required,email,lte=100" validate:"required"`
	Password string `json:"password" binding:"required,gte=6,lte=20" validate:"required"`
}

func (user UserType) ReturnUserString() string {
	switch user {
	case USER:
		return string(USER)
	case RIDER:
		return string(RIDER)
	}
	return string(UNKNOWN)
}

func (u *User) BuildNewUser(user UserDTO) (*User, error) {
	if err := compareBothPasswords(user.Pass, user.ConfirmPassword); err != nil {
		return &User{}, err
	}
	Pass, err := hashPassword(user.Pass)
	if err != nil {
		return &User{}, err
	}
	return &User{
		Email: user.Email,
		Password: Pass,
		UserType: UserType(user.UserType),
	}, nil
}

func (user *User) ComparePassword(dbpass, pass string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(dbpass), []byte(pass)); err != nil {
		return errors.New("incorrect log in credentials")
	}
	return nil
}

func hashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost);
	if err != nil {
		return "", err
	}
	password = string(hashPassword)
	return password, nil
}

func compareBothPasswords(password, confirmPassword string) error {
	if password != confirmPassword {
		return fmt.Errorf("incorrect passwords")
	}
	return nil
}
