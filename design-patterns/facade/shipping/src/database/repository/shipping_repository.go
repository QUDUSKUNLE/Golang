package repository

import (
	"github.com/google/uuid"
	"github.com/QUDUSKUNLE/shipping/src/model"
)

func (database *Database) QueryCreateShipping(shipping model.Shipping) error {
	query := `INSERT INTO shippings (id, user_id, product_description, pick_up_address, delivery_address, product_type) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := database.Exec(query, shipping.ID, shipping.UserID, shipping.Description, shipping.PickUpAddress, shipping.DeliveryAddress, shipping.ProductType)
	if err != nil {
		return err
	}
	return nil
}

func (database *Database) QueryShippings(userID uuid.UUID, status string) ([]model.Shipping, error) {
	shippings := []model.Shipping{}
	query := `SELECT * FROM shippings WHERE id=$1`
	if err := database.Select(&shippings, query, userID); err != nil {
		return shippings, err
	}
	return shippings, nil
}
