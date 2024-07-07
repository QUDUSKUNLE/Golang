package shipping

import (
	"fmt"

	"github.com/QUDUSKUNLE/shipping/src/ledger"
	"github.com/QUDUSKUNLE/shipping/src/model"
	"github.com/QUDUSKUNLE/shipping/src/notification"
)

type UserAdaptor struct {
	user *model.User
	userRepository *ledger.UserRepository
	notification *notification.Notification
}

func NewUserAdaptor() *UserAdaptor {
	return &UserAdaptor{
		user: &model.User{},
		userRepository: &ledger.UserRepository{},
		notification: &notification.Notification{},
	}
} 

func (userAdaptor *UserAdaptor) RegisterNewUser(user model.UserDTO) error {
	fmt.Println("Start a new user registration")
	buildUser, err := userAdaptor.user.BuildNewUser(user)
	if err != nil {
		return err
	}
	// Save use in the database
	if err := userAdaptor.userRepository.UserLedger(buildUser); err != nil {
		return err
	}
	userAdaptor.notification.SendRegistrationNotification()
	fmt.Println("Successfully registered a new user")
	return nil
}
