package pointer

import (
	"fmt"
	"errors"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

type Stringer interface {
	String() string
}

type Wallet struct {
	// It's private outside
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	(*w).balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return (*w).balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

/*
Suppose you have a variable
age := 20
Using &age get you the pointer to the variable, its memory address
When you have the pointer to the variable, you can get the value it points
to by using the * operator
age := 20
ageptr = &age
agevalue = *ageptr
*/

func Increment(val *int) *int {
	*val = *val + 1
	return val
}
