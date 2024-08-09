package services

import (
	"fmt"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
)

func (httpHandler *InternalServicesHandler) SaveUser(userDto domain.UserDTO) error {
	fmt.Println("Initiate a new user registration")

	servicesHandler := httpHandler.NewInternalServicesFacade()
	buildUser, err := servicesHandler.userService.BuildNewUser(userDto)
	if err != nil {
		return err
	}

	// Save use in the database
	err = httpHandler.internal.SaveUserAdaptor(*buildUser);
	if err != nil {
		return err
	}
	servicesHandler.notificationService.SendRegistrationNotification()
	fmt.Println("Successfully registered a new user")
	return nil
}

func (httpHandler *InternalServicesHandler) ResetPassword(userDto domain.ResetPasswordDto) error {
	fmt.Println("Initiate a reset password")

	_, err := httpHandler.internal.ReadUserByEmailAdaptor(userDto.Email)
	if err != nil {
		return fmt.Errorf("user %s with the email: %s", err.Error(), userDto.Email)
	}
	servicesHandler := httpHandler.NewInternalServicesFacade()
	servicesHandler.notificationService.SendRegistrationNotification()
	fmt.Println("Reset Password link sent successfully")
	return nil
}
