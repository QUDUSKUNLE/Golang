package ledger

import (
	"fmt"
	"github.com/google/uuid"
)

type DeliveryLedger struct {}

func (ledger *DeliveryLedger) Ledger(accountID uuid.UUID, product string) (shippingID string, err error) {
	fmt.Printf("Make a Delivery ledger entry for accountID %s with productType %s.\n", accountID, product)
	return "1", nil
}
