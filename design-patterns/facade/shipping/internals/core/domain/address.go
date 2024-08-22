package domain

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Address struct {
	Email             string  `json:"email" validate:"required,email"`
	Province          string  `json:"province"`
	TerminalAddressID string  `json:"terminal_address_id"`
	Description       string  `json:"description" validate:"min=3,max=20,required"`
	FirstName         string  `json:"first_name" validate:"required"`
	LastName          string  `json:"last_name" validate:"max=50,required"`
	StreetNo          string  `json:"street_no" validate:"max=50,required"`
	StreetName        string  `json:"street_name" validate:"max=50,required"`
	City              string  `json:"city" validate:"max=50,required"`
	State             string  `json:"state" validate:"max=50,required"`
	PhoneNo           string  `json:"phone_no" validate:"max=50,required"`
	Zip               string  `json:"zip" validate:"max=50,required"`
	Country           Country `json:"country" validate:"max=50,required"`
}

func (a Address) Value() (driver.Value, error) {
	// Serialize the Address struct into a format suitable for storage
	// For example, you might serialize it into a JSON string
	addressJSON, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}
	return string(addressJSON), nil
}

func (a *Address) Scan(value interface{}) error {
	addressJSON, ok := value.(string)
	if !ok {
		return errors.New("unexpected type for address")
	}
	return json.Unmarshal([]byte(addressJSON), a)
}
