package shipping

import "fmt"

type ProductType string

const (
	Animal ProductType = "Animal"
	Plant ProductType	 = "Plant"
	Appareal ProductType = "Appareal"
	Book ProductType   = "Book"
	Cosmetics ProductType = "Cosmetics"
	Electronics ProductType = "Electronics"
	Watery  ProductType = "Watery"
	Ammunition ProductType = "Ammunition"
	Unknown ProductType = "Unknown"
)

func (product ProductType) PrintProduct() string {
	switch product {
	case Animal:
		return string(Animal)
	case Plant:
		return string(Plant)
	case Appareal:
		return string(Appareal)
	case Book:
		return string(Book)
	case Cosmetics:
		return string(Cosmetics)
	case Electronics:
		return string(Electronics)
	case Watery:
		return string(Watery)
	case Ammunition:
		return string(Ammunition)
	}
	return string(Unknown)
}

type Product struct {
	productType ProductType
}

func NewProduct(product ProductType) *Product {
	return &Product{
		productType: product,
	}
}

func (p *Product) CheckProduct(product ProductType) error {
	if p.productType.PrintProduct() != string(product) {
		return fmt.Errorf("wrong product is submitted: %s", string(product))
	}
	return nil
}
