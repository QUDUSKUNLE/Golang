package domain

type ProductType string

const (
	Animal      ProductType = "Animal"
	Plant       ProductType = "Plant"
	Appareal    ProductType = "Appareal"
	Book        ProductType = "Book"
	Cosmetics   ProductType = "Cosmetics"
	Electronics ProductType = "Electronics"
	Watery      ProductType = "Watery"
	Ammunition  ProductType = "Ammunition"
	Unknown     ProductType = "Unknown"
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
