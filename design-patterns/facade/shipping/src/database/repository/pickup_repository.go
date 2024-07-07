package repository

import (
	"github.com/QUDUSKUNLE/shipping/src/model"
)

func (database *Database) QueryCreatePickUp(pickUp model.PickUp) error {
	query := `INSERT INTO pickups (id, shipping_id, carrier_id, pick_up_at, status) VALUES ($1, $2, $3, $4, $5)`
	_, err := database.Exec(query, pickUp.ID, pickUp.ShippingID, pickUp.CarrierID, pickUp.PickUpAt, pickUp.Status)
	if err != nil {
		return err
	}
	return nil
}

func (database *Database) QueryUpdatePickUp(pickUp model.PickUp) error {
	query := `UPDATE pickups SET pick_up_at=$3, updated_at=$4, status=$5 WHERE id=$1 AND shipping_id=$2 AND carrier_id=$6`
	_, err := database.Exec(query, pickUp.ID, pickUp.ShippingID, pickUp.PickUpAt, pickUp.UpdatedAt, pickUp.Status, pickUp.CarrierID)
	if err != nil {
		return err
	}
	return nil
}
