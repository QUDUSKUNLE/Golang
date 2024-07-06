package handlers

import (
	"net/http"

	"github.com/QUDUSKUNLE/shipping/src"
	"github.com/QUDUSKUNLE/shipping/src/dto"
	"github.com/labstack/echo/v4"
)

func Register(context echo.Context) error {
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
	userAdaptor := shipping.NewUserAdaptor()

	err := userAdaptor.RegisterNewUser(*user);
	if err != nil {
		if err.Error() == "user`s already exist" {
			return context.JSON(http.StatusConflict, map[string]string{"message": "User already registered", "success": "false" })
		}
		if err.Error() == `incorrect passwords` {
			return context.JSON(http.StatusBadRequest, map[string]string{"message": err.Error(), "success": "false" })
		}
		return context.JSON(http.StatusNotAcceptable, map[string]string{"message": err.Error(), "success": "false" })
	}
	// Process valid user data
	return context.JSON(http.StatusOK, map[string]string{
		"message": "User registered successfully",
		"success": "true",
	})
}


func Login(context echo.Context) error {
	loginDto := new(dto.LogInDTO)
	// Bind loginDto
	if err := context.Bind(loginDto); err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}
	// Validate user input
	if err := context.Validate(loginDto); err != nil {
		return err
	}
	// Initiate a new login adaptor
	login := shipping.NewLogInAdaptor()

   token, err := login.LoginUser(*loginDto);
	 if err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
			"success": "false",
		})
	}
	// Process valid user data
	return context.JSON(http.StatusOK, map[string]string{
		"message": "User login successfully",
		"token": token,
		"success": "true",
	})
}
