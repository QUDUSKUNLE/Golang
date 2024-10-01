package facade

import "fmt"

type Wallet struct {
	balance int
}

func NewWallet() *Wallet {
	return &Wallet{
		balance: 0,
	}
}

func (wallet *Wallet) CreditWallet(amount int) {
	wallet.balance += amount
	fmt.Println("Wallet balance added successfully")
}

func (wallet *Wallet) DebitWallet(amount int) error {
	if wallet.balance < amount {
		return fmt.Errorf("insufficient balance")
	}
	wallet.balance -= amount
	fmt.Println("Wallet debited successfully")
	return nil
}

