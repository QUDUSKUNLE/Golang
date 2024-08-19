package domain

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	USER    UserType = "USER"
	CARRIER UserType = "CARRIER"
	ADMIN 	UserType = "ADMIN"
)

type (
	User struct {
		ID        uuid.UUID       `json:"id" gorm:"uuid;primaryKey"`
		CreatedAt *time.Time      `json:"created_at"`
		UpdatedAt *time.Time      `json:"updated_at"`
		DeletedAt *gorm.DeletedAt `gorm:"index" json:"-"`
		Email     string          `json:"email" gorm:"unique"`
		Password  string          `json:"password"`
		UserType  UserType        `json:"user_type"`

		Shippings  []Shipping  `json:"shippings" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
		Addresses  []Location  `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
		Packagings []Packaging `json:"packagings" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
		Parcels    []Parcel    `json:"parcels" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	}
	UserDto struct {
		Email           string   `json:"email" binding:"required,email,lte=100" validate:"required,email"`
		Password        string   `json:"password" binding:"required,gte=6,lte=20" validate:"required,gte=6,lte=20"`
		ConfirmPassword string   `json:"confirm_password" binding:"required,gte=6,lte=20" validate:"required,gte=6,lte=20"`
		UserType        UserType `json:"user_type" binding:"required" validate:"required"`
	}
	LogInDto struct {
		Email    string `json:"email" binding:"required,email,lte=100" validate:"required"`
		Password string `json:"password" binding:"required,gte=6,lte=20" validate:"required"`
	}
	ResetPasswordDto struct {
		Email string `json:"email" binding:"required" validate:"required,email"`
	}
	UserType string
	Response struct {
		Result interface{} `json:"result"`
		Success bool	`json:"success"`
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
	return nil
}

func (user *User) AfterCreate(scope *gorm.DB) error {
	if user.UserType == CARRIER {
		scope.Model(&Carrier{}).Create(&Carrier{UserID: user.ID})
	}
	return nil
}

func (u *User) BuildNewUser(user UserDto) (*User, error) {
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
