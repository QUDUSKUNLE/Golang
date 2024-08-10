package integration

type(
	METHOD string
	PACKAGE_TYPE string
	ENDPOINTS string
)

const (
	// METHODS
	GET METHOD = "GET"
	PUT METHOD = "PUT"
	POST METHOD = "POST"

	// ENDPOINTS
  RATES ENDPOINTS = "rates"
	SHIPMENTS ENDPOINTS = "shipments"
	PACKAGING ENDPOINTS = "packaging"
	PICKUP ENDPOINTS = "pickups"
	ADDRESSES ENDPOINTS = "addresses"

	// PACKAGE_TYPE
	BOX PACKAGE_TYPE = "box"
	ENVELOPE PACKAGE_TYPE = "envelope"
	SOFT_PACKAGING PACKAGE_TYPE = "soft_packaging"
)

func (method METHOD) PrintMethod() string {
	switch method {
	case GET:
		return string(GET)
	case POST:
		return string(POST)
	case PUT:
		return string(PUT)
	}
	return "Unknown"
}

func (endpoint ENDPOINTS) PrintEndpoint() string {
	switch endpoint {
	case RATES:
		return string(RATES)
	case SHIPMENTS:
		return string(SHIPMENTS)
	case PACKAGING:
		return string(PACKAGING)
	case PICKUP:
		return string(PICKUP)
	case ADDRESSES:
		return string(ADDRESSES)
	}
	return "Unknown"
}
