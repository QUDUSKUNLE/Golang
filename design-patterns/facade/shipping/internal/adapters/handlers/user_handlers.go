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
		return handler.ComputeErrorResponse(http.StatusBadRequest, err.Error(), context)
	}

	err := handler.servicesAdapter.SaveUser(*user);
	if err != nil {
		if err.Error() == string(USER_ALREADY_EXIST) {
			return handler.ComputeErrorResponse(http.StatusConflict, USER_ALREADY_REGISTERED, context)
		}

		if err.Error() == string(INCORRECT_PASSWORDS) {
			return handler.ComputeErrorResponse(http.StatusBadRequest, err.Error(), context)
		}
		return handler.ComputeErrorResponse(http.StatusConflict, USER_ALREADY_REGISTERED, context)
	}
	// Process valid user data
	return handler.ComputeResponseMessage(http.StatusOK,USER_REGISTERED_SUCCESSFULLY, context)
}

func (handler *HTTPHandler) Login(context echo.Context) error {
	loginDto := new(domain.LogInDTO)
	if err := handler.ValidateStruct(context, loginDto); err != nil {
		return handler.ComputeErrorResponse(http.StatusBadRequest, err.Error(), context)
	}
	// Initiate a new login adaptor
	user, err := handler.servicesAdapter.LogInUserAdaptor(*loginDto)
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusBadRequest, err.Error(), context)
	}
	Token, err := handler.GenerateAccessToken(CurrentUser{ID: user.ID, UserType: string(user.UserType)})
	if err != nil {
		return handler.ComputeErrorResponse(http.StatusBadRequest, err.Error(), context)
	}
	// Process valid user data
	return handler.ComputeResponseMessage(http.StatusOK, Token, context)
}

func (handler *HTTPHandler) Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	ID := claims.ID
	return handler.ComputeResponseMessage(http.StatusOK, ID.String(), c)
}

func (handler *HTTPHandler) ResetPassword(context echo.Context) error {
	resetPasswordDto := new(domain.ResetPasswordDto)
	if err := handler.ValidateStruct(context, resetPasswordDto); err != nil {
		return handler.ComputeErrorResponse(http.StatusBadRequest, err.Error(), context)
	}
	if err := handler.servicesAdapter.ResetPassword(*resetPasswordDto); err != nil {
		return handler.ComputeErrorResponse(http.StatusBadRequest, err.Error(), context)
	}
	return handler.ComputeResponseMessage(http.StatusOK, RESET_EMAIL_SENT_SUCCESSFULLY, context)
}
