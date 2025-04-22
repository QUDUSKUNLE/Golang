package dto

import (
	"errors"

	"github.com/QUDUSKUNLE/microservices/gateway/db"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
)

type (
	UserDto struct {
		Email           string      `json:"email" validate:"email,required"`
		Password        string      `json:"password" validate:"min=8,required"`
		ConfirmPassword string      `json:"confirm_password" validate:"eqfield=Password"`
		UserType        db.UserEnum `json:"user_type"`
	}
	LogInDto struct {
		Email    string `json:"email" validate:"email,required"`
		Password string `json:"password" validate:"min=8,required"`
	}
	UpdateNinDto struct {
		Nin    string `json:"nin" validate:"required"`
		UserID string `json:"user_id" validate:""`
	}
	ResetPasswordDto struct {
		Email string `json:"email" validate:"email,required"`
	}
	CurrentUser struct {
		ID       string `json:"id"`
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
	Local struct {
		*db.User
	}
)

func BuildNewUser(user UserDto) (*db.User, error) {
	Password, err := HashPassword(user.Password)
	if err != nil {
		return &db.User{}, err
	}
	return &db.User{
		Email:    pgtype.Text{String: user.Email, Valid: true},
		Password: Password,
		UserType: user.UserType,
	}, nil
}

func ComparePassword(user db.User, pass string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass)); err != nil {
		return errors.New("incorrect log in credentials")
	}
	return nil
}

func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	password = string(hashPassword)
	return password, nil
}
