package repository

import (
	"github.com/QUDUSKUNLE/shipping/src/model"
)

func (database *Database) QueryCreateShipping(shipping *model.Shipping) error {
	query := `INSERT INTO shippings (id, user_id, product_description, pick_up_address, delivery_address, product_type) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := database.Exec(query, shipping.ID, shipping.UserID, shipping.Description, shipping.PickUpAddress, shipping.DeliveryAddress, shipping.ProductType)
	if err != nil {
		return err
	}
	return nil
}
