package handlers

import (
	"net/http"

	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func (handler *HTTPHandler) Register(context echo.Context) error {
	user := new(domain.UserDTO)
	if err := handler.ValidateStruct(context, user); err != nil {
		return err
	}

	err := handler.ServicesAdapter.SaveUser(*user);
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
	if err := handler.ValidateStruct(context, loginDto); err != nil {
		return err
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

func (handler *HTTPHandler) ResetPassword(context echo.Context) error {
	resetPasswordDto := new(domain.ResetPasswordDto)
	if err := handler.ValidateStruct(context, resetPasswordDto); err != nil {
		return context.JSON(http.StatusBadRequest, echo.Map{"Success": false, "Message": err.Error()})
	}
	if err := handler.ServicesAdapter.ResetPassword(*resetPasswordDto); err != nil {
		return context.JSON(http.StatusBadRequest, echo.Map{"Success": false, "Message": err.Error()})
	}
	return context.JSON(http.StatusOK, echo.Map{
		"Message": "Reset email sent successfully.",
		"Success": true,
	})
}
