package domain

type (
	ProductType string
  PACKAGE_TYPE string
	WEIGHT_UNIT string
	ITEM_TYPE string
	Currency string
	CASH_ON_DELIVERY string
	SHIPMENT_PURPOSE string
)

const (
	Animal      ProductType = "Animal"
	Plant       ProductType = "Plant"
	Appareal    ProductType = "Appareal"
	Book        ProductType = "Book"
	Cosmetics   ProductType = "Cosmetics"
	Electronics ProductType = "Electronics"
	Watery      ProductType = "Watery"
	Ammunition  ProductType = "Ammunition"

	// PACKAGE_TYPE
	box            PACKAGE_TYPE = "box"
	envelope       PACKAGE_TYPE = "envelope"
	soft_packaging PACKAGE_TYPE = "soft_packaging"

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
	USD Currency = "USD"
	TZS Currency = "TZS"
	UGX Currency = "UGX"
	ZAR Currency = "ZAR"

	// Cash on delivery
	False CASH_ON_DELIVERY = "false"
	True CASH_ON_DELIVERY = "true"

	// Shipment Purpose
	COMMERCIAL SHIPMENT_PURPOSE = "commercial"
	PERSONAL SHIPMENT_PURPOSE = "personal"
	SAMPLE    SHIPMENT_PURPOSE = "sample"
	RETURN_AFTER_REPAIR SHIPMENT_PURPOSE = "return-after-repair"
	RETURN_FOR_REPAIR SHIPMENT_PURPOSE = "return-for-repair"
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
	case box:
		return string(box)
	case envelope:
		return string(envelope)
	case soft_packaging:
		return string(soft_packaging)
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
	case False:
		return string(False)
	case True:
		return string(True)
	}
	return "Unknown"
}

func (ship SHIPMENT_PURPOSE) PrintShipmentPurpose() string {
	switch ship {
	case COMMERCIAL:
		return string(COMMERCIAL)
	case PERSONAL:
		return string(PERSONAL)
	case SAMPLE:
		return string(SAMPLE)
	case RETURN_AFTER_REPAIR:
		return string(RETURN_AFTER_REPAIR)
	case RETURN_FOR_REPAIR:
		return string(RETURN_FOR_REPAIR)
	}
	return "Unknown"
}
