package shipping

import (
	"fmt"

	"github.com/QUDUSKUNLE/shipping/src/account"
	"github.com/QUDUSKUNLE/shipping/src/ledger"
	"github.com/QUDUSKUNLE/shipping/src/notification"
)

type User struct {
	user *account.User
	userLedger *ledger.UserLedger
	notification *notification.Notification
}

func NewUser(accountID string) *User {
	return &User{
		user: account.RegisterNewUser(accountID),
		userLedger: &ledger.UserLedger{},
		notification: &notification.Notification{},
	}
} 

func (u *User) NewUser(email, password, confirmPassword string) error {
	fmt.Println("Start a new user registration")
	// Compare both passwords
	if err := u.user.CompareBothPasswords(password, confirmPassword); err != nil {
		return err
	}

	// Check if email is not registered before
	if err := u.user.CheckEmail(email); err != nil {
		return err
	}

	// Hash user password, this should be done at the database level
	password, err := u.user.HashPassword(password);
	if err != nil {
		return err;
	}

	// Save use in the database
	_, err = u.userLedger.Ledger(password)
	if err != nil {
		return err
	}
	u.notification.SendRegistrationNotification()
	fmt.Println("Successfully registered a new user")
	return nil
}
