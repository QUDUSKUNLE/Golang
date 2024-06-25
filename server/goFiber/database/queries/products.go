package queries

import (
	"gofiber/models"
	"github.com/jmoiron/sqlx"
)

type Fiber struct {
	*sqlx.DB
}

func (query *Fiber) GetProducts() ([]models.Product, error) {
	products := []models.Product{}
	q := `SELECT * from products`
	if err := query.Select(&products, q); err != nil {
		return products, err
	}
	return products, nil
}

func (query *Fiber) GetProduct(id int) (models.Product, error) {
	product := models.Product{}
	q := `SELECT * FROM products WHERE id = $1`
	if err := query.Get(&product, q, id); err != nil {
		return product, err
	}
	return product, nil
}

func (query *Fiber) CreateProduct(product *models.Product) error {
	q := `INSERT INTO products VALUES ($1, $2, $3, $4, $5)`
	_, err := query.Exec(q, product.ID, product.Amount, product.Name, product.Description, product.Category)
	if err != nil {
		return err;
	}
	return nil
}

func (query *Fiber) UpdateProduct(id int, p *models.Product) error {
	q := `UPDATE products SET name=$2, description=$3, category=$4, amount=$5 WHERE id=$1`;
	_, err := query.Exec(q, id, p.Name, p.Description, p.Category, p.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (query *Fiber) DeleteProduct(id int) error {
	q := `DELETE FROM products WHERE id=$1`
	_, err := query.Exec(q, id);
	if err != nil {
		return err
	}
	return nil
}
