package services

import (
	"fmt"

	"github.com/QUDUSKUNLE/shipping/internal/core/ledger"
	"github.com/QUDUSKUNLE/shipping/internal/core/model"
)

type UserAdaptor struct {
	userService *model.User
	userRepositoryService *ledger.UserRepository
	notificationService *Notification
}

func NewUserAdaptor(userDto model.UserDTO) error {
	fmt.Println("Initiate a new user registration")
	userAdaptor := &UserAdaptor{
		userService: &model.User{},
		userRepositoryService: &ledger.UserRepository{},
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
