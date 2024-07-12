package services

import (
	"fmt"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
)

type UserAdaptor struct {
	userService *domain.User
	notificationService *Notification
}


func (httpHandler *ServicesHandler) SaveUserAdaptor(userDto domain.UserDTO) error {
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
	err = httpHandler.Internal.SaveUser(*buildUser);
	if err != nil {
		return err
	}
	userAdaptor.notificationService.SendRegistrationNotification()
	fmt.Println("Successfully registered a new user")
	return nil
}
