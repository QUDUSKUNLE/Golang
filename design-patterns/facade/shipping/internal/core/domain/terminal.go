package domain

import "fmt"

type PackagingDto struct {
	Height      float32      `json:"height" binding:"required" validate:"required"`
	Length      float32      `json:"length" binding:"required" validate:"required"`
	Name        string       `json:"name" binding:"required" validate:"required"`
	Size_Unit   string       `json:"size_unit" binding:"required" validate:"required"`
	Type        PACKAGE_TYPE `json:"type" binding:"required" validate:"required"`
	Width       float32      `json:"width" binding:"required" validate:"required"`
	Weight      float32      `json:"weight" binding:"required" validate:"required"`
	Weight_Unit string       `json:"weight_unit" binding:"required" validate:"required"`
}

type TerminalAddressDto struct {
	Address
	Is_residential *bool `json:"Is_residential" binding:"required" validate:"required"`
}

func (packaging *PackagingDto) BuildNewPackaging(pack PackagingDto) map[string]interface{} {
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

func (address *PackagingDto) BuildNewAddress(addr Address) map[string]interface{} {
	return map[string]interface{}{
		"city": addr.City,
		"country": addr.Country.PrintCountry(),
		"email": addr.Email,
		"first_name": addr.FirstName,
		"is_residential": true,
		"last_name": addr.LastName,
		"line1": addr.StreetNo,
		"line2": addr.StreetName,
		"name": fmt.Sprintf("%s %s", addr.FirstName, addr.LastName),
		"phone": addr.PhoneNo,
		"state": addr.State,
		"zip": addr.Zip,
	}
}
