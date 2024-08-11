package services

import (
	"fmt"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
)

func (internalHandler *InternalServicesHandler) SaveUser(userDto domain.UserDto) error {
	servicesHandler := internalHandler.NewInternalServicesFacade()
	buildUser, err := servicesHandler.userService.BuildNewUser(userDto)
	if err != nil {
		return err
	}
	// Save use in the database
	err = internalHandler.internal.SaveUserAdaptor(*buildUser);
	if err != nil {
		return err
	}
	servicesHandler.notificationService.SendRegistrationNotification()
	return nil
}

func (internalHandler *InternalServicesHandler) ResetPassword(userDto domain.ResetPasswordDto) error {
	_, err := internalHandler.internal.ReadUserByEmailAdaptor(userDto.Email)
	if err != nil {
		return fmt.Errorf("user %s with the email: %s", err.Error(), userDto.Email)
	}
	servicesHandler := internalHandler.NewInternalServicesFacade()
	servicesHandler.notificationService.SendResetPasswordNotification()
	return nil
}
