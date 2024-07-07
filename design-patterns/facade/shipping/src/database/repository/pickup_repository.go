package repository


import (
	"github.com/QUDUSKUNLE/shipping/src/model"
)

func (database *Database) QueryCreatePickUp(shipping model.PickUp) error {
	query := `INSERT INTO pickups (id, shipping_id, carrier_id, pick_up_at, status) VALUES ($1, $2, $3, $4, $5)`
	_, err := database.Exec(query, shipping.ID, shipping.ShippingID, shipping.CarrierID, shipping.PickUpAt, shipping.Status)
	if err != nil {
		return err
	}
	return nil
}
