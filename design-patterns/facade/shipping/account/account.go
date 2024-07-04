package shipping

import "fmt"


type User struct {
	ID 		   string
	Email 	 string
	Password string
}

func NewAccount(accountID string) *User {
	return &User{
		ID: accountID,
	}
}

func (ship *User) CheckAccount(accountID string) error {
	if ship.ID != accountID {
		return fmt.Errorf("accountID %s is not known", accountID)
	}
	return nil
}

func (account *User) CheckEmail(Email string) error {
	if account.Email != Email {
		return fmt.Errorf("email %s is not known", Email)
	}
	return nil
}
