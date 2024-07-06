package model

import (
	"errors"
	"fmt"
	"time"

	"github.com/QUDUSKUNLE/shipping/src/dto"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserType string

const (
	USER UserType = "USER"
	RIDER UserType = "RIDER"
	UNKNOWN UserType = "UNKNOWN"
)

type User struct {
	ID 		   		uuid.UUID
	Email 	 		string
	Password 		string
	UserType    UserType
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}

func NewUser(ID uuid.UUID) *User {
	return &User{
		ID: ID,
	}
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

func BuildUser(user dto.UserDTO) (*User, error) {
	if err := compareBothPasswords(user.Password, user.ConfirmPassword); err != nil {
		return &User{}, err
	}
	Pass, err := hashPassword(user.Password)
	if err != nil {
		return &User{}, err
	}
	return &User{
		Email: user.Email,
		Password: Pass,
		UserType: UserType(user.UserType),
		CreatedAt: time.Now(),
	}, nil
}

func (user *User) CheckUser(userID uuid.UUID) error {
	if user.ID != userID {
		return fmt.Errorf("accountID %s is not known", userID)
	}
	return nil
}

func (user *User) ComparePassword(pass string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(pass), []byte(user.Password)); err != nil {
		return errors.New("incorrect password")
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
