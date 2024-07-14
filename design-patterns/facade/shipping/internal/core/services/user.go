package services

import (
	"fmt"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
)

type UserAdaptor struct {
	userService *domain.User
	notificationService *Notification
}

func (httpHandler *ServicesHandler) SaveUser(userDto domain.UserDTO) error {
	fmt.Println("Initiate a new user registration")
	userAdaptor := &UserAdaptor{
		userService: &domain.User{},
		notificationService: &Notification{},
	}
	buildUser, err := userAdaptor.userService.BuildNewUser(userDto)
	if err != nil {
		return err
	}
	// Save use in the database
	err = httpHandler.Internal.SaveUserAdaptor(*buildUser);
	if err != nil {
		return err
	}
	userAdaptor.notificationService.SendRegistrationNotification()
	fmt.Println("Successfully registered a new user")
	return nil
}

func (httpHandler *ServicesHandler) ResetPassword(userDto domain.ResetPasswordDto) error {
	fmt.Println("Initiate a reset password")
	userAdaptor := &UserAdaptor{
		notificationService: &Notification{},
	}

	_, err := httpHandler.Internal.ReadUserByEmailAdaptor(userDto.Email)
	if err != nil {
		return fmt.Errorf("user %s with the email: %s", err.Error(), userDto.Email)
	}
	userAdaptor.notificationService.SendRegistrationNotification()
	fmt.Println("Reset Password link sent successfully")
	return nil
}
