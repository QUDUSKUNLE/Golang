package handlers

import (
	"net/http"

	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func (handler *HTTPHandler) Register(context echo.Context) error {
	user := new(domain.UserDTO)
	// Bind userDto
	if err := context.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Validate user input
	if err := context.Validate(user); err != nil {
		return context.JSON(http.StatusBadRequest, echo.Map{"Success": false, "Message": err.Error()})
	}

	err := handler.ServicesAdapter.SaveUserAdaptor(*user);
	if err != nil {
		if err.Error() == "user`s already exist" {
			return context.JSON(http.StatusConflict, echo.Map{"Message": "User already registered", "Success": false })
		}
		if err.Error() == `incorrect passwords` {
			return context.JSON(http.StatusBadRequest, echo.Map{"Message": err.Error(), "Success": false })
		}
		return context.JSON(http.StatusNotAcceptable, echo.Map{"Message": err.Error(), "Success": false })
	}
	// Process valid user data
	return context.JSON(http.StatusOK, echo.Map{
		"Message": "User registered successfully",
		"Success": true,
	})
}

func (handler *HTTPHandler) Login(context echo.Context) error {
	loginDto := new(domain.LogInDTO)
	// Bind loginDto
	if err := context.Bind(loginDto); err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}
	// Validate user input
	if err := context.Validate(loginDto); err != nil {
		return context.JSON(http.StatusBadRequest, echo.Map{"Success": false, "Message": err.Error() })
	}
	// Initiate a new login adaptor
	user, err := handler.ServicesAdapter.LogInUserAdaptor(*loginDto)
	if err != nil {
		return context.JSON(http.StatusBadRequest, echo.Map{
			"Message": err.Error(),
			"Success": "false",
		})
	}
	token, err := handler.GenerateAccessToken(CurrentUser{ID: user.ID, UserType: string(user.UserType)})
	if err != nil {
		return context.JSON(http.StatusBadRequest, echo.Map{
			"Message": err.Error(),
			"Success": "false",
		})
	}
	// Process valid user data
	return context.JSON(http.StatusOK, echo.Map{"Token": token})
}

func (handler *HTTPHandler) Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	ID := claims.ID
	return c.JSON(http.StatusOK, echo.Map{"Message": ID.String()})
}
