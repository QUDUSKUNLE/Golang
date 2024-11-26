package domain

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

type Country string

const (
	NG      Country = "NG"
	US      Country = "US"
	UK      Country = "UK"
	UAE     Country = "UAE"
	NIGERIA Country = "NG"
)

func (country Country) PrintCountry() string {
	switch country {
	case US:
		return string(US)
	case UK:
		return string(UK)
	case UAE:
		return string(UAE)
	case NG:
		return string(NG)
	}
	return "Unknown"
}

type (
	Address struct {
		Email             string  `json:"email" validate:"email,required"`
		Province          string  `json:"province"`
		TerminalAddressID string  `json:"terminal_address_id"`
		Description       string  `json:"description" validate:"min=3,max=100,required"`
		FirstName         string  `json:"first_name" validate:"min=3,max=50,required"`
		LastName          string  `json:"last_name" validate:"min=3,max=50,required"`
		StreetNo          string  `json:"street_no" validate:"gte=1,lt=1000,required"`
		StreetName        string  `json:"street_name" validate:"min=3,max=50,required"`
		City              string  `json:"city" validate:"min=2,max=50,required"`
		State             string  `json:"state" validate:"min=2,max=50,required"`
		PhoneNo           string  `json:"phone_no" validate:"max=50,required"`
		Zip               string  `json:"zip" validate:"max=50,required"`
		Country           Country `json:"country" validate:"max=50,required"`
	}
	AddressParamDto struct {
		AddressID uuid.UUID `param:"address_id" validate:"uuid,required"`
	}
	AddressQueryDto struct {
		Description string `query:"description"`
	}
	Contact struct {
		PhoneNumbers []string `json:"phone_numbers" validate:"gt=0,dive,required"`
		WhatsApps    string   `json:"whatsApp" validate:"required"`
		Twitter      string   `json:"twitter" validate:"required"`
	}
)

func (address Address) Value() (driver.Value, error) {
	// Serialize the Address struct into a format suitable for storage
	// For example, you might serialize it into a JSON string
	addressJson, err := json.Marshal(address)
	if err != nil {
		return nil, err
	}
	return string(addressJson), nil
}

func (address *Address) Scan(value interface{}) error {
	addressJson, ok := value.(string)
	if !ok {
		return errors.New("unexpected type for address")
	}
	return json.Unmarshal([]byte(addressJson), address)
}

func (a Contact) Value() (driver.Value, error) {
	// Serialize the Address struct into a format suitable for storage
	// For example, you might serialize it into a JSON string
	contactJSON, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}
	return string(contactJSON), nil
}

func (a *Contact) Scan(value interface{}) error {
	contactJSON, ok := value.(string)
	if !ok {
		return errors.New("unexpected type for address")
	}
	return json.Unmarshal([]byte(contactJSON), a)
}
