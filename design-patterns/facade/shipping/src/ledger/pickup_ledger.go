package ledger

import (
	"fmt"
	"github.com/QUDUSKUNLE/shipping/src/model"
)

type PickUpLedger struct {}

func (ledger *PickUpLedger) NewLedger(pick model.PickUp) (shippingID string, err error) {
	fmt.Printf("Make pick up ledger entry for accountID %s with productType %s.\n", pick.ID, pick.Status)
	return "1", nil
}


func (ledger *PickUpLedger) UpdateLedger(update model.PickUp) (shippingID string, err error) {
	fmt.Printf("Make pick up ledger entry for accountID %s with productType %s.\n", update.ID, update.Status)
	return "1", nil
}
