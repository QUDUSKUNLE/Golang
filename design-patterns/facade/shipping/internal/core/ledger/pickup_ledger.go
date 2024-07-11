package ledger

import (
	"github.com/QUDUSKUNLE/shipping/internal/adapters/repository"
	"github.com/QUDUSKUNLE/shipping/internal/core/model"
)

type PickUpRepository struct {}

func (ledger *PickUpRepository) NewLedger(pick model.PickUp) error {
	// Open database conection
	db, err := repository.OpenDBConnection()
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
	db, err := repository.OpenDBConnection()
	if err != nil {
		return err
	}
	if err := db.QueryUpdatePickUp(update); err != nil {
		return err
	}
	return nil
}
