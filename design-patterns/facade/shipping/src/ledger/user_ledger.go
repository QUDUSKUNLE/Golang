package ledger

import (
	"fmt"
)

type UserLedger struct {}

func (ledger *UserLedger) RegisterLedger(accountID string) error {
	fmt.Printf("Register a new user %s\n", accountID)
	return nil
}
