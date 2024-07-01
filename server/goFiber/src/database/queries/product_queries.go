package queries

import (
	"database/sql"
	"github.com/QUDUSKUNLE/gofiber/src/models"
	"github.com/jmoiron/sqlx"
)

type Fiber struct {
	*sqlx.DB
}

func (query *Fiber) QueryGetProducts() ([]models.Product, error) {
	products := []models.Product{}
	q := `SELECT * from products`
	err := query.Select(&products, q);
	if err != sql.ErrNoRows {
		return []models.Product{}, err
	}
	return products, nil
}

func (query *Fiber) QueryGetProduct(id int) (*models.Product, error) {
	product := models.Product{}
	q := `SELECT * FROM products WHERE id = $1`
	if err := query.Get(&product, q, id); err != nil {
		return &models.Product{}, err
	}
	return &product, nil
}

func (query *Fiber) QueryCreateProduct(product *models.Product) error {
	q := `INSERT INTO products VALUES ($1, $2, $3, $4, $5)`
	_, err := query.Exec(q, product.ID, product.Amount, product.Name, product.Description, product.Category)
	if err != nil {
		return err;
	}
	return nil
}

func (query *Fiber) QueryUpdateProduct(id int, p *models.Product) error {
	q := `UPDATE products SET name=$2, description=$3, category=$4, amount=$5 WHERE id=$1`;
	_, err := query.Exec(q, id, p.Name, p.Description, p.Category, p.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (query *Fiber) QueryDeleteProduct(id int) error {
	q := `DELETE FROM products WHERE id=$1`
	_, err := query.Exec(q, id);
	if err != nil {
		return err
	}
	return nil
}
