package models

import (
	"sync"
)

type Product struct {
	sync.Mutex 
	ID          int    `json:"id"`
	Name 				string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Amount      int    `json:"amount"`
}

var product Product

func (p *Product) Unique() (id int) {
	p.Lock()
  defer p.Unlock()
	if p.ID == 0 {
		p.ID = 1
	}
  id = p.ID
  p.ID++
  return
}

func NewProduct() *Product {
	return &Product{
			ID: product.Unique(),
	}
}
