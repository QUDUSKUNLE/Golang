package handlers

import (
	"net/http"
	"github.com/QUDUSKUNLE/shipping/internals/core/domain"
	"github.com/labstack/echo/v4"
)

// @Summary Register a new user
// @Description register a new user
// @Tags User
// @Accept json
// @Produce json
// @Param body body domain.UserDto true "Register a user"
// @Failure 409 {object} domain.Response
// @Success 201 {object} domain.Response
// @Router /register [post]
func (handler *HTTPHandler) Register(context echo.Context) error {
	user := new(domain.UserDto)
	if err := ValidateStruct(context, user); err != nil {
		return ComputeErrorResponse(http.StatusBadRequest, err, context)
	}
	err := handler.internalServicesAdapter.SaveUser(*user);

	if err != nil {
		if err.Error() == string(USER_ALREADY_EXIST) {
			return ComputeErrorResponse(http.StatusConflict, USER_ALREADY_REGISTERED, context)
		}
		return ComputeErrorResponse(http.StatusConflict, USER_ALREADY_REGISTERED, context)
	}
	// Process valid user data
	return ComputeResponseMessage(http.StatusCreated, USER_REGISTERED_SUCCESSFULLY, context)
}

// @Summary Sign in a user
// @Description sign in a user
// @Tags User
// @Accept json
// @Produce json
// @Param body body domain.LogInDto true "Sign in a user"
// @failure 400 {object} domain.Response
// @Success 200 {object} domain.Response
// @Router /login [post]
func (handler *HTTPHandler) Login(context echo.Context) error {
	loginDto := new(domain.LogInDto)
	if err := ValidateStruct(context, loginDto); err != nil {
		return ComputeErrorResponse(http.StatusBadRequest, err, context)
	}
	// Initiate a new login adaptor
	user, err := handler.internalServicesAdapter.LogInUserAdaptor(*loginDto)
	if err != nil {
		return ComputeErrorResponse(http.StatusNotFound, err.Error(), context)
	}
	Token, err := GenerateAccessToken(CurrentUser{ID: user.ID, UserType: string(user.UserType)})
	if err != nil {
		return ComputeErrorResponse(http.StatusBadRequest, err.Error(), context)
	}
	// Process valid user data
	return ComputeResponseMessage(http.StatusOK, Token, context)
}

// @Summary Reset a user password
// @Description Reset a user password
// @Tags User
// @Accept json
// @Produce json
// @Param body body domain.ResetPasswordDto true "Reset password"
// @failure 400 {object} domain.Response
// @Success 200 {object} domain.Response
// @Router /resetpassword [post]
func (handler *HTTPHandler) ResetPassword(context echo.Context) error {
	resetPasswordDto := new(domain.ResetPasswordDto)
	if err := ValidateStruct(context, resetPasswordDto); err != nil {
		return ComputeErrorResponse(http.StatusBadRequest, err.Error(), context)
	}
	if err := handler.internalServicesAdapter.ResetPassword(*resetPasswordDto); err != nil {
		return ComputeErrorResponse(http.StatusBadRequest, err.Error(), context)
	}
	return ComputeResponseMessage(http.StatusOK, RESET_EMAIL_SENT_SUCCESSFULLY, context)
}
