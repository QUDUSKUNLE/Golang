package integration

type METHOD string
type PACKAGE_TYPE string

const (
	// METHODS
	GET string = "GET"
	PUT string = "PUT"
	POST string = "POST"

	// ENDPOINTS
  RATES string = "rates"
	SHIPMENT string = "shipments"
	PACKAGING string = "packaging"
	PICKUP string = "pickups"
	ADDRESSES string = "addresses"

	// PACKAGE_TYPE
	BOX PACKAGE_TYPE = "box"
	ENVELOPE PACKAGE_TYPE = "envelope"
	SOFT_PACKAGING PACKAGE_TYPE = "soft_packaging"
)
