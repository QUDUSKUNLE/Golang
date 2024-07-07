package handlers

import (
	"net/http"

	"github.com/QUDUSKUNLE/shipping/src"
	"github.com/QUDUSKUNLE/shipping/src/model"
	"github.com/QUDUSKUNLE/shipping/src/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func Register(context echo.Context) error {
	user := new(model.UserDTO)
	// Bind userDto
	if err := context.Bind(user); err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}
	// Validate user input
	if err := context.Validate(user); err != nil {
		return err
	}

	err := shipping.NewUserAdaptor(*user);
	if err != nil {
		if err.Error() == "user`s already exist" {
			return context.JSON(http.StatusConflict, echo.Map{"message": "User already registered", "success": "false" })
		}
		if err.Error() == `incorrect passwords` {
			return context.JSON(http.StatusBadRequest, echo.Map{"message": err.Error(), "success": "false" })
		}
		return context.JSON(http.StatusNotAcceptable, echo.Map{"message": err.Error(), "success": "false" })
	}
	// Process valid user data
	return context.JSON(http.StatusOK, echo.Map{
		"message": "User registered successfully",
		"success": "true",
	})
}

func Login(context echo.Context) error {
	loginDto := new(model.LogInDTO)
	// Bind loginDto
	if err := context.Bind(loginDto); err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}
	// Validate user input
	if err := context.Validate(loginDto); err != nil {
		return context.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": err.Error() })
	}
	// Initiate a new login adaptor
	token, err := shipping.NewLogInAdaptor(*loginDto)
	 if err != nil {
		return context.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
			"success": "false",
		})
	}
	// Process valid user data
	return context.JSON(http.StatusOK, echo.Map{"token": token})
}

func Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*utils.JwtCustomClaims)
	ID := claims.ID
	return c.JSON(http.StatusOK, echo.Map{"message": ID.String()})
}
