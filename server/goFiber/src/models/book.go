package models

import (
	"time"
	"errors"
	"database/sql/driver"
	"encoding/json"
	"github.com/google/uuid"
)

type Book struct {
	ID 				uuid.UUID   `json:"id" xml:"id" form:"id" validate:"required,uuid"`
	CreatedAt  time.Time 	`json:"created_at" xml:"created_at" form:"created_at" db:"created_at"`
  UpdatedAt  time.Time 	`json:"updated_at" xml:"updated_at" form:"updated_at"  db:"updated_at"`
	Title     string      `json:"title" xml:"title" form:"title" validate:"required,lte=255"`
	Author		string      `json:"author" xml:"author" form:"author" validate:"required,lte=255"`
	BookStatus int       	`json:"book_status" xml:"book_status" form:"book_status" validate:"required,len=1" db:"book_status"`
  BookAttrs  BookAttrs 	`json:"book_attrs" xml:"book_attrs" form:"book_attrs" validate:"required,dive" db:"book_attrs"`
}

// BookAttrs struct to describe book attributes.
type BookAttrs struct {
	Picture     string `json:"picture" xml:"picture" form:"picture"`
	Description string `json:"description" xml:"description" form:"description"`
	Rating      int    `json:"rating" xml:"rating" form:"rating" validate:"min=1,max=10"`
}

// Value make the BookAttrs struct implement the driver.Valuer interface.
// This method simply returns the JSON-encoded representation of the struct.
func (b BookAttrs) Value() (driver.Value, error) {
	return json.Marshal(b)
}

// Scan make the BookAttrs struct implement the sql.Scanner interface.
// This method simply decodes a JSON-encoded value into the struct fields.
func (b *BookAttrs) Scan(value interface{}) error {
	j, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(j, &b)
}

func NewBook() *Book {
	return &Book{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
