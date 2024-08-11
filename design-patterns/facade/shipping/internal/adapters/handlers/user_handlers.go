package handlers

import (
	"net/http"

	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// @Summary Register a new user
// @Description Register a new user
// @Tags User
// @Accept json
// @Produce json
// @Param Request body domain.UserDto true "Register a user"
// @failure 409 {object} domain.RegisterResponse
// @Success 201 {object} domain.RegisterResponse
// @Router /register [post]
func (handler *HTTPHandler) Register(context echo.Context) error {
	user := new(domain.UserDto)
	if err := handler.ValidateStruct(context, user); err != nil {
		return handler.ComputeErrorResponse(http.StatusBadRequest, err, context)
	}
	err := handler.internalServicesAdapter.SaveUser(*user);
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
	return handler.ComputeResponseMessage(http.StatusCreated,USER_REGISTERED_SUCCESSFULLY, context)
}

// @Summary Sign in a user
// @Description Sign in a user
// @Tags User
// @Accept json
// @Produce json
// @Param Request body domain.LogInDto true "Sign in a user"
// @failure 400 {object} domain.RegisterResponse
// @Success 200 {object} domain.RegisterResponse
// @Router /login [post]
func (handler *HTTPHandler) Login(context echo.Context) error {
	loginDto := new(domain.LogInDto)
	if err := handler.ValidateStruct(context, loginDto); err != nil {
		return handler.ComputeErrorResponse(http.StatusBadRequest, err, context)
	}
	// Initiate a new login adaptor
	user, err := handler.internalServicesAdapter.LogInUserAdaptor(*loginDto)
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

// @Summary Reset a user password
// @Description Reset a user password
// @Tags User
// @Accept json
// @Produce json
// @Param Request body domain.ResetPasswordDto true "Reset password"
// @failure 400 {object} domain.RegisterResponse
// @Success 200 {object} domain.RegisterResponse
// @Router /reset_password [post]
func (handler *HTTPHandler) ResetPassword(context echo.Context) error {
	resetPasswordDto := new(domain.ResetPasswordDto)
	if err := handler.ValidateStruct(context, resetPasswordDto); err != nil {
		return handler.ComputeErrorResponse(http.StatusBadRequest, err.Error(), context)
	}
	if err := handler.internalServicesAdapter.ResetPassword(*resetPasswordDto); err != nil {
		return handler.ComputeErrorResponse(http.StatusBadRequest, err.Error(), context)
	}
	return handler.ComputeResponseMessage(http.StatusOK, RESET_EMAIL_SENT_SUCCESSFULLY, context)
}
