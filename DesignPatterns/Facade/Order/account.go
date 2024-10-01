package facade

import "fmt"

type Account struct {
	name string
}

func NewAccount(accoutName string) *Account {
	return &Account{
		name: accoutName,
	}
}

func (acc *Account) CheckAccount(accountName string) error {
	if acc.name != accountName {
		return fmt.Errorf("Account Name is incorrect")
	}
	fmt.Println("Account verified")
	return nil
}
