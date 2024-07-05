package handlers

import (
	"net/http"

	"github.com/QUDUSKUNLE/shipping/src"
	"github.com/QUDUSKUNLE/shipping/src/dto"
	"github.com/labstack/echo/v4"
)

func NewUser(context echo.Context) error {
	user := new(dto.UserDTO)
	// Bind userDto
	if err := context.Bind(user); err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}
	// Validate user input
	if err := context.Validate(user); err != nil {
		return err
	}
	// Initiate a new user registration
	newUser := shipping.NewUserAdaptor(user.Email)

	if err := newUser.NewUser(user.Email, user.Password, user.ConfirmPassword); err != nil {
		return context.JSON(http.StatusNotAcceptable, map[string]string{"message": err.Error(), "success": "false" })
	}
	// Process valid user data
	return context.JSON(http.StatusOK, map[string]string{
		"message": "User registered successfully",
		"success": "true",
	})
}

