package domain

import "fmt"

type (
	Terminal             struct{}
	TerminalPackagingDto struct {
		Packagings []SingleTerminalPackagingDto `json:"Packagings" binding:"required" validate:"required,dive,required"`
	}
	TerminalParcelDto struct {
		Parcels []SingleTerminalParcelDto `json:"Parcels" binding:"required" validate:"required,dive,required"`
	}
	TerminalParcelItemDto struct {
		Description string    `json:"description" validate:"required"`
		HS_CODE     string    `json:"hs_code"`
		Name        string    `json:"name" validate:"required"`
		Type        ITEM_TYPE `json:"type" validate:"required"`
		Currency    Currency  `json:"currency" validate:"required"`
		Value       float32   `json:"value" validate:"required"`
		Quantity    int       `json:"quantity" validate:"required"`
		Weight      float32   `json:"weight" validate:"required"`
	}
	SingleTerminalPackagingDto struct {
		Height      float32      `json:"height" binding:"required" validate:"required"`
		Length      float32      `json:"length" binding:"required" validate:"required"`
		Name        string       `json:"name" binding:"required" validate:"required"`
		Size_Unit   string       `json:"size_unit" binding:"required" validate:"required"`
		Type        PACKAGE_TYPE `json:"type" binding:"required" validate:"required"`
		Width       float32      `json:"width" binding:"required" validate:"required"`
		Weight      float32      `json:"weight" binding:"required" validate:"required"`
		Weight_Unit string       `json:"weight_unit" binding:"required" validate:"required"`
	}
	SingleTerminalParcelDto struct {
		Description       string                  `json:"description" validate:"required"`
		Items             []TerminalParcelItemDto `json:"items" binding:"required" validate:"required,dive,required"`
		Metadata          map[string]interface{}  `json:"metadata"`
		Packaging         string                  `json:"packaging" validate:"required"`
		Proof_Of_Payments []string                `json:"proof_of_payments" binding:"required" validate:"required,dive,required"`
		Rec_docs          []string                `json:"rec_docs" binding:"required" validate:"required,dive,required"`
		Weight_unit       WEIGHT_UNIT             `json:"weight_unit" binding:"required" validate:"required"`
	}
	TerminalRatesQueryDto struct {
		Currency          Currency         `json:"currency"`
		PickUpAddressID   string           `json:"pickup_address_id"`
		DeliveryAddressID string           `json:"delivery_address_id"`
		ShipmentID        *string          `json:"shipment_id"`
		ParcelID          string           `json:"parcel_id"`
		CashOnDelivery    CASH_ON_DELIVERY `json:"cash_on_delivery"`
	}
)

func (packaging *Terminal) BuildNewTerminalPackaging(pack SingleTerminalPackagingDto) map[string]interface{} {
	return map[string]interface{}{
		"height":      pack.Height,
		"length":      pack.Length,
		"name":        pack.Name,
		"size_unit":   pack.Size_Unit,
		"type":        pack.Type.PrintPackageType(),
		"width":       pack.Width,
		"weight":      pack.Weight,
		"weight_unit": pack.Weight_Unit,
	}
}

func (address *Terminal) BuildNewTerminalAddress(addr Address) map[string]interface{} {
	return map[string]interface{}{
		"city":           addr.City,
		"country":        addr.Country.PrintCountry(),
		"email":          addr.Email,
		"first_name":     addr.FirstName,
		"is_residential": true,
		"last_name":      addr.LastName,
		"line1":          addr.StreetNo,
		"line2":          addr.StreetName,
		"name":           fmt.Sprintf("%s %s", addr.FirstName, addr.LastName),
		"phone":          addr.PhoneNo,
		"state":          addr.State,
		"zip":            addr.Zip,
	}
}

func (parcel *Terminal) BuildNewTerminalParcel(parse SingleTerminalParcelDto) map[string]interface{} {
	return map[string]interface{}{
		"description":       parse.Description,
		"items":             buildNewTerminalParcelItem(parse.Items),
		"metadata":          parse.Metadata,
		"packaging":         parse.Packaging,
		"proof_of_payments": parse.Proof_Of_Payments,
		"rec_docs":          parse.Rec_docs,
		"weight_unit":       parse.Weight_unit,
	}
}

func buildNewTerminalParcelItem(parse []TerminalParcelItemDto) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, parcelItem := range parse {
		result = append(result, map[string]interface{}{
			"description": parcelItem.Description,
			"hs_code":     parcelItem.HS_CODE,
			"name":        parcelItem.Name,
			"type":        parcelItem.Type,
			"currency":    parcelItem.Currency.PrintCurrency(),
			"value":       parcelItem.Value,
			"quantity":    parcelItem.Quantity,
			"weight":      parcelItem.Weight,
		})
	}
	return result
}

func (rates *Terminal) BuildNewTerminalRatesQuery(query TerminalRatesQueryDto) string {
	return fmt.Sprintf("currency=%s&pickup_address=%s&delivery_address=%s&cash_on_delivery=%s&parcel_id=%s", query.Currency.PrintCurrency(), query.PickUpAddressID, query.DeliveryAddressID, query.CashOnDelivery.PrintCashOnDelivery(), query.ParcelID)
}