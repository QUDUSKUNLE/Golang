package shipping

import (
	"fmt"

	"github.com/QUDUSKUNLE/shipping/src/ledger"
	"github.com/QUDUSKUNLE/shipping/src/model"
	"github.com/QUDUSKUNLE/shipping/src/notification"
)

type UserAdaptor struct {
	userService *model.User
	userRepositoryService *ledger.UserRepository
	notificationService *notification.Notification
}

func NewUserAdaptor(userDto model.UserDTO) error {
	fmt.Println("Initiate a new user registration")
	userAdaptor := &UserAdaptor{
		userService: &model.User{},
		userRepositoryService: &ledger.UserRepository{},
		notificationService: &notification.Notification{},
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

func UsersAdaptor() ([]model.User, error) {
	userAdaptor := &UserAdaptor{
		userRepositoryService: &ledger.UserRepository{},
	}
	// Save use in the database
	users, err := userAdaptor.userRepositoryService.QueryUsers();
	if err != nil {
		return []model.User{}, err
	}
	return users, nil
}
