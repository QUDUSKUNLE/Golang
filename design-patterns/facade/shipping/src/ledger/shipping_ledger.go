package ledger

import (
	"github.com/QUDUSKUNLE/shipping/src/database"
	"github.com/QUDUSKUNLE/shipping/src/model"
)

type ScheduleShippingLedger struct {}

func (ledger *ScheduleShippingLedger) ShippingLedger(shipping *model.Shipping) error {
	// Open database conection
	db, err := database.OpenDBConnection()
	if err != nil {
		return err
	}
	if err := db.QueryCreateShipping(shipping); err != nil {
		return err
	}
	return nil
}
