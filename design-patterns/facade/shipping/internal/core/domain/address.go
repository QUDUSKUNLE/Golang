package domain

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Address struct {
	Email             string `json:"email" validate:"required,email"`
	Province          string `json:"province"`
	TerminalAddressID string `json:"terminal_address_id"`
	Description       string `json:"description"`
	FirstName         string  `json:"first_name" validate:"required"`
	LastName          string  `json:"last_name" binding:"required,max=50" validate:"required"`
	StreetNo          string  `json:"street_no" binding:"required,max=50" validate:"required"`
	StreetName        string  `json:"street_name" binding:"required,max=50" validate:"required"`
	City              string  `json:"city" binding:"required,max=50" validate:"required"`
	State             string  `json:"state" binding:"required,max=50" validate:"required"`
	PhoneNo           string  `json:"phone_no" binding:"required,max=50" validate:"required"`
	Zip               string  `json:"zip" binding:"required,max=50" validate:"required"`
	Country           Country `json:"country" binding:"required,max=50" validate:"required"`
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
