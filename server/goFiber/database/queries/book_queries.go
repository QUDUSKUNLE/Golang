package queries

import (
	"database/sql"
	"gofiber/models"

	"github.com/google/uuid"
)

func (query *Fiber) QueryGetBooks() ([]models.Book, error) {
	books := []models.Book{}
	q := `SELECT * FROM books`
	err := query.Select(&books, q);
	if err != sql.ErrNoRows {
		return books, err
	}
	return books, nil
}

func (query *Fiber) QueryGetBook(id uuid.UUID) (models.Book, error) {
	book := models.Book{}
	q := `SELECT * FROM books WHERE id = $1`
	if err := query.Get(&book, q, id); err != nil {
		return book, err
	}
	return book, nil
}

func (query *Fiber) QueryCreateBook(book *models.Book) error {
	q := `INSERT INTO books VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := query.Query(q, book.ID, book.CreatedAt, book.UpdatedAt, book.Title, book.Author, book.BookStatus, book.BookAttrs)
	if err != nil {
		return err;
	}
	return nil
}

func (query *Fiber) QueryUpdateBook(id uuid.UUID, book *models.Book) error {
	q := `UPDATE books SET updated_at=$2, title=$3, author=$4, book_status=$5, book_attrs=$6 WHERE id=$1`;
	_, err := query.Exec(q, id, book.UpdatedAt, book.Title, book.Author, book.BookStatus, book.BookAttrs)
	if err != nil {
		return err
	}
	return nil
}

func (query *Fiber) QueryDeleteBook(id uuid.UUID) error {
	q := `DELETE FROM books WHERE id=$1`
	_, err := query.Exec(q, id);
	if err != nil {
		return err
	}
	return nil
}
