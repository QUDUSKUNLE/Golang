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
	ID 		   						uuid.UUID 	`json:"ID" gorm:"uuid;primaryKey"`
	Email 	 						string    	`json:"Email" gorm:"unique"`
	Password						string    	`json:"Password"`
	UserType    				UserType  	`json:"UserType"`
	CreatedAt 					time.Time 	`json:"CreatedAt"`
	UpdatedAt 					*time.Time 	`json:"UpdatedAt,omitempty"`
	Shippings   				[]Shipping 	`json:"Shippings" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	PickUps    					[]PickUp  	`json:"PickUps" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type UserDTO struct {
	Email    						string 			`json:"Email" binding:"required,email,lte=100" validate:"required"`
	Password 				    string 			`json:"Password" binding:"required,gte=6,lte=20" validate:"required"`
	ConfirmPassword 		string 			`json:"ConfirmPassword" binding:"required,gte=6,lte=20" validate:"required"`
	UserType 						string 			`json:"UserType" binding:"required" validate:"required"`
}

type LogInDTO struct {
	Email  							string 			`json:"Email" binding:"required,email,lte=100" validate:"required"`
	Password  					string 			`json:"Password" binding:"required,gte=6,lte=20" validate:"required"`
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

func (user *User) BeforeCreate(scope *gorm.DB) error {
	user.ID = uuid.New()
	return nil
}

func (u *User) BuildNewUser(user UserDTO) (*User, error) {
	if err := compareBothPasswords(user.Password, user.ConfirmPassword); err != nil {
		return &User{}, err
	}
	Password, err := hashPassword(user.Password)
	if err != nil {
		return &User{}, err
	}
	return &User{
		Email: user.Email,
		Password: Password,
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
