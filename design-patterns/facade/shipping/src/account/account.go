package account

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID 		   		string
	Email 	 		string
	Password 		string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}

func NewUser(userID string) *User {
	return &User{
		ID: userID,
	}
}

func RegisterNewUser(email string) *User {
	return &User{
		Email: email,
	}
}

func (user *User) CheckUser(userID string) error {
	if user.ID != userID {
		return fmt.Errorf("accountID %s is not known", userID)
	}
	return nil
}

func (user *User) CheckEmail(Email string) error {
	if user.Email != Email {
		return fmt.Errorf("email %s is not known", Email)
	}
	return nil
}

func (user *User) ComparePassword(pass string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(pass), []byte(user.Password)); err != nil {
		return errors.New("incorrect password")
	}
	return nil
}

func (user *User) HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost);
	if err != nil {
		return "", err
	}
	password = string(hashPassword)
	return password, nil
}

func (user *User) CompareBothPasswords(password, confirmPassword string) error {
	if password != confirmPassword {
		return fmt.Errorf("incorrect passwords")
	}
	return nil
}