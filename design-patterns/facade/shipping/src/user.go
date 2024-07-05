package shipping

import (
	"fmt"

	"github.com/QUDUSKUNLE/shipping/src/model"
	"github.com/QUDUSKUNLE/shipping/src/ledger"
	"github.com/QUDUSKUNLE/shipping/src/notification"
)

type UserAdaptor struct {
	user *model.User
	userLedger *ledger.UserLedger
	notification *notification.Notification
}

func NewUserAdaptor(accountID string) *UserAdaptor {
	return &UserAdaptor{
		user: model.RegisterNewUser(accountID),
		userLedger: &ledger.UserLedger{},
		notification: &notification.Notification{},
	}
} 

func (userAdaptor *UserAdaptor) NewUser(email, password, confirmPassword string) error {
	fmt.Println("Start a new user registration")
	// Compare both passwords
	if err := userAdaptor.user.CompareBothPasswords(password, confirmPassword); err != nil {
		return err
	}
	// Check if email is not registered before
	if err := userAdaptor.user.CheckEmail(email); err != nil {
		return err
	}
	// Hash user password, this should be done at the database level
	password, err := userAdaptor.user.HashPassword(password);
	if err != nil {
		return err;
	}
	// Save use in the database
	if err := userAdaptor.userLedger.RegisterLedger(password); err != nil {
		return err
	}
	userAdaptor.notification.SendRegistrationNotification()
	fmt.Println("Successfully registered a new user")
	return nil
}
