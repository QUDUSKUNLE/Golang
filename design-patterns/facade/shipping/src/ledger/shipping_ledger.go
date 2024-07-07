package ledger

import (
	"github.com/QUDUSKUNLE/shipping/src/database"
	"github.com/QUDUSKUNLE/shipping/src/model"
)

type ShippingRepository struct {}

func (ledger *ShippingRepository) ShippingLedger(shipping model.Shipping) error {
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
