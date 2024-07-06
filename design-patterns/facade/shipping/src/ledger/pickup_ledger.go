package ledger

import (
	"fmt"
	"github.com/google/uuid"
)

type PickUpLedger struct {}

func (ledger *PickUpLedger) Ledger(accountID uuid.UUID, product string) (shippingID string, err error) {
	fmt.Printf("Make pick up ledger entry for accountID %s with productType %s.\n", accountID, product)
	return "1", nil
}
