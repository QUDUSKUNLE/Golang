package domain

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
	USER    	UserType = "USER"
	CARRIER   UserType = "CARRIER"
	UNKNOWN 	UserType = "UNKNOWN"
)

type User struct {
	gorm.Model
	ID        uuid.UUID  `json:"ID" gorm:"uuid;primaryKey"`
	Email     string     `json:"Email" gorm:"unique"`
	Password  string     `json:"Password"`
	UserType  UserType   `json:"UserType"`
	CreatedAt time.Time  `json:"CreatedAt"`
	UpdatedAt time.Time  `json:"UpdatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Shippings []Shipping `json:"Shippings" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	Addresses []Location `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type UserDTO struct {
	Email           string `json:"Email" binding:"required,email,lte=100" validate:"required"`
	Password        string `json:"Password" binding:"required,gte=6,lte=20" validate:"required"`
	ConfirmPassword string `json:"ConfirmPassword" binding:"required,gte=6,lte=20" validate:"required"`
	UserType        string `json:"UserType" binding:"required" validate:"required"`
}

type LogInDTO struct {
	Email    string `json:"Email" binding:"required,email,lte=100" validate:"required"`
	Password string `json:"Password" binding:"required,gte=6,lte=20" validate:"required"`
}

type ResetPasswordDto struct {
	Email string `json:"Email" binding:"required" validate:"required,email"`
}

func (user UserType) ReturnUserString() string {
	switch user {
	case USER:
		return string(USER)
	case CARRIER:
		return string(CARRIER)
	}
	return string(UNKNOWN)
}

func (user *User) BeforeCreate(scope *gorm.DB) error {
	user.ID = uuid.New()
	return nil
}

func (user *User) AfterCreate(scope *gorm.DB) error {
	if user.UserType == CARRIER {
		scope.Model(&Carrier{}).Create(&Carrier{UserID: user.ID})
	}
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
		Email:    user.Email,
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
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
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
