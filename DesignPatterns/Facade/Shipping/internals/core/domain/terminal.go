package domain

import (
	"fmt"
	"github.com/google/uuid"
)

type (
	Terminal             struct{}
	TerminalPackagingDto struct {
		Packagings []SingleTerminalPackagingDto `json:"Packagings" validate:"gt=0,dive,required"`
	}
	TerminalParcelDto struct {
		Parcels []SingleTerminalParcelDto `json:"Parcels" validate:"gt=0,dive,required"`
	}
	TerminalShipmentDto struct {
		Shipments []SingleTerminalShipmentDto `json:"shipments" validate:"gt=0,dive,required"`
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
	TerminalRatesQueryDto struct {
		Currency          Currency         `json:"currency" validate:"required"`
		PickUpAddressID   string           `json:"pickup_address_id"`
		DeliveryAddressID string           `json:"delivery_address_id"`
		ShipmentID        *string          `json:"shipment_id"`
		ParcelID          string           `json:"parcel_id"`
		CashOnDelivery    CASH_ON_DELIVERY `json:"cash_on_delivery" validate:"required"`
	}
	SingleTerminalShipmentDto struct {
		PickUpAddressID   string           `json:"pickup_address_id"`
		DeliveryAddressID string           `json:"delivery_address_id"`
		Parcels           []string         `json:"parcel_id" validate:"gt=0,dive,required"`
		ReturnAddressID   string           `json:"return_address_id"`
		CarrierID         uuid.UUID        `json:"carrier_id" validate:"uuid,required"`
		ProductType       ProductType      `json:"product_type" validate:"required"`
		Description       string           `json:"description" validate:"required,gte=6,lte=100"`
		ShipmentPurpose   SHIPMENT_PURPOSE `json:"shipment_purpose" validate:"required"`
		ShipmentType      CASH_ON_DELIVERY `json:"shipment_type"`
	}
	SingleTerminalPackagingDto struct {
		Height      float32      `json:"height" validate:"required"`
		Length      float32      `json:"length" validate:"required"`
		Name        string       `json:"name" validate:"required"`
		Size_Unit   string       `json:"size_unit" validate:"required"`
		Type        PACKAGE_TYPE `json:"type" validate:"required"`
		Width       float32      `json:"width" validate:"required"`
		Weight      float32      `json:"weight" validate:"required"`
		Weight_Unit string       `json:"weight_unit" validate:"required"`
	}
	SingleTerminalParcelDto struct {
		Description       string                  `json:"description" validate:"required"`
		Items             []TerminalParcelItemDto `json:"items" validate:"required,dive,required"`
		Metadata          map[string]interface{}  `json:"metadata"`
		Packaging         string                  `json:"packaging" validate:"required"`
		Proof_Of_Payments []string                `json:"proof_of_payments" validate:"required,dive,required"`
		Rec_docs          []string                `json:"rec_docs" validate:"required,dive,required"`
		Weight_unit       WEIGHT_UNIT             `json:"weight_unit" validate:"required"`
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

func (ship *Terminal) BuildNewTerminalShipment(shipment SingleTerminalShipmentDto) map[string]interface{} {
	return map[string]interface{}{
		"address_from": shipment.PickUpAddressID,
		"address_to":   shipment.DeliveryAddressID,
		"metadata":     "",
		// "parcel":           shipment.Parcels[0],
		"address_return":   shipment.PickUpAddressID,
		"shipment_purpose": shipment.ShipmentPurpose.PrintShipmentPurpose(),
		"parcels":          shipment.Parcels,
		"shipment_type":    shipment.ShipmentType.PrintCashOnDelivery(),
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
