package ledger

import (
	"github.com/QUDUSKUNLE/shipping/internal/adapters/repository"
	"github.com/QUDUSKUNLE/shipping/internal/core/model"
	"github.com/google/uuid"
)

type ShippingRepository struct{}

func (ledger *ShippingRepository) ShippingLedger(shipping model.Shipping) error {
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

func (ledger *ShippingRepository) QueryShippingLedger(userID uuid.UUID, status string) ([]model.Shipping, error) {
	// Open repository conection
	db, err := repository.OpenDBConnection()
	if err != nil {
		return []model.Shipping{}, err
	}
	shippings, err := db.QueryShippings(userID, status)
	if err != nil {
		return shippings, err
	}
	return shippings, nil
}
