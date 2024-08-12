package domain

type (
	ProductType string
  PACKAGE_TYPE string
	WEIGHT_UNIT string
	ITEM_TYPE string
	Currency string
	CASH_ON_DELIVERY string
)

const (
	Animal      ProductType = "animal"
	Plant       ProductType = "plant"
	Appareal    ProductType = "appareal"
	Book        ProductType = "book"
	Cosmetics   ProductType = "cosmetics"
	Electronics ProductType = "electronics"
	Watery      ProductType = "watery"
	Ammunition  ProductType = "ammunition"

	// PACKAGE_TYPE
	BOX            PACKAGE_TYPE = "box"
	ENVELOPE       PACKAGE_TYPE = "envelope"
	SOFT_PACKAGING PACKAGE_TYPE = "soft_packaging"

	// WEIGHT_UNIT
	KG WEIGHT_UNIT = "kg"

	// ITEM_TYPE
	DOCUMENT ITEM_TYPE = "document"
	PARCEL ITEM_TYPE = "parcel"

	// Currency
	AED Currency = "AED"
	AUD Currency = "AUD"
	CAD Currency = "CAD"
	CNY Currency = "CNY"
	EUR Currency = "EUR"
	GBP Currency = "GBP"
	GHS Currency = "GHS"
	HKD Currency = "HKD"
	KES Currency = "KES"
	NGN Currency = "NGN"
	TZS Currency = "TZS"
	UGX Currency = "UGX"
	USD Currency = "USD"
	ZAR Currency = "ZAR"

	// Cash on delivery
	false CASH_ON_DELIVERY = "false"
	true CASH_ON_DELIVERY = "true"
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
	return "Unknown"
}

func (packageType PACKAGE_TYPE) PrintPackageType() string {
	switch packageType {
	case BOX:
		return string(BOX)
	case ENVELOPE:
		return string(ENVELOPE)
	case SOFT_PACKAGING:
		return string(SOFT_PACKAGING)
	}
	return "Unknown"
}

func (currency Currency) PrintCurrency() string {
	switch currency {
	case NGN:
		return string(NGN)
	case USD:
		return string(USD)
	}
	return "Unknown"
}


func (catch CASH_ON_DELIVERY) PrintCashOnDelivery() string {
	switch catch {
	case false:
		return string(false)
	case true:
		return string(true)
	}
	return "Unknown"
}
