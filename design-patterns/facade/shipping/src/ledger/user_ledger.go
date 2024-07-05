package ledger

import (
	"fmt"
)

type UserLedger struct {}

func (ledger *UserLedger) Ledger(accountID string) (userID string, err error) {
	fmt.Printf("Register a new user %s\n", accountID)
	return accountID, nil
}
