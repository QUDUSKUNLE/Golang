package ledger

import (
	"github.com/QUDUSKUNLE/shipping/internal/adapters/repository"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
)

func (ledger *Ledger) NewLedger(pick domain.PickUp) error {
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


func (ledger *Ledger) UpdateLedger(update domain.PickUp) error {
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
