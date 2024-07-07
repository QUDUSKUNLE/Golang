package ledger

import (
	"github.com/QUDUSKUNLE/shipping/src/database"
	"github.com/QUDUSKUNLE/shipping/src/model"
)

type PickUpRepository struct {}

func (ledger *PickUpRepository) NewLedger(pick model.PickUp) error {
	// Open database conection
	db, err := database.OpenDBConnection()
	if err != nil {
		return err
	}
	if err := db.QueryCreatePickUp(pick); err != nil {
		return err
	}
	return nil
}


func (ledger *PickUpRepository) UpdateLedger(update model.PickUp) error {
	// Open database conection
	db, err := database.OpenDBConnection()
	if err != nil {
		return err
	}
	if err := db.QueryUpdatePickUp(update); err != nil {
		return err
	}
	return nil
}
