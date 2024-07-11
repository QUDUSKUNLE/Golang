package ledger

import (
	"github.com/QUDUSKUNLE/shipping/internal/adapters/repository"
	"github.com/QUDUSKUNLE/shipping/internal/core/domain"
	"github.com/google/uuid"
)

func (ledger *Ledger) ShippingLedger(shipping domain.Shipping) error {
	// Open repository conection
	db, err := repository.OpenDBConnection()
	if err != nil {
		return err
	}
	if err := db.QueryCreateShipping(shipping); err != nil {
		return err
	}
	return nil
}

func (ledger *Ledger) QueryShippingLedger(userID uuid.UUID, status string) ([]domain.Shipping, error) {
	// Open repository conection
	db, err := repository.OpenDBConnection()
	if err != nil {
		return []domain.Shipping{}, err
	}
	shippings, err := db.QueryShippings(userID, status)
	if err != nil {
		return shippings, err
	}
	return shippings, nil
}
