package ledger

import (
	"fmt"
)

type PickUpLedger struct {}

func (ledger *PickUpLedger) Ledger(accountID, product string) (shippingID string, err error) {
	fmt.Printf("Make pick up ledger entry for accountID %s with productType %s.\n", accountID, product)
	return "1", nil
}
