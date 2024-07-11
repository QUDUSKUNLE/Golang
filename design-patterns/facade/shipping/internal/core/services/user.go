package services

import (
	"fmt"

	"github.com/QUDUSKUNLE/shipping/internal/core/ledger"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
)

type UserAdaptor struct {
	userService *domain.User
	userRepositoryService *ledger.Ledger
	notificationService *Notification
}

func NewUserAdaptor(userDto domain.UserDTO) error {
	fmt.Println("Initiate a new user registration")
	userAdaptor := &UserAdaptor{
		userService: &domain.User{},
		userRepositoryService: &ledger.Ledger{},
		notificationService: &Notification{},
	}
	buildUser, err := userAdaptor.userService.BuildNewUser(userDto)
	if err != nil {
		return err
	}
	// Save use in the database
	err = userAdaptor.userRepositoryService.QueryCreateUser(buildUser);
	if err != nil {
		return err
	}
	userAdaptor.notificationService.SendRegistrationNotification()
	fmt.Println("Successfully registered a new user")
	return nil
}

func (httpHandler *ServicesHandler) SaveUserAdaptor(userDto domain.UserDTO) error {
	fmt.Println("Initiate a new user registration")
	userAdaptor := &UserAdaptor{
		userService: &domain.User{},
		userRepositoryService: &ledger.Ledger{},
		notificationService: &Notification{},
	}
	buildUser, err := userAdaptor.userService.BuildNewUser(userDto)
	if err != nil {
		return err
	}
	// Save use in the database
	err = userAdaptor.userRepositoryService.QueryCreateUser(buildUser);
	if err != nil {
		return err
	}
	userAdaptor.notificationService.SendRegistrationNotification()
	fmt.Println("Successfully registered a new user")
	return nil
}
