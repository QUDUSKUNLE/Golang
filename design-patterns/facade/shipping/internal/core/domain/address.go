package domain

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Address struct {
	Email             string `json:"Email" validate:"required,email"`
	Province          string `json:"Province"`
	ExternalAddressID string `json:"ExternalAddressID"`
	Description       string `json:"Description"`
	FirstName         string  `json:"FirstName" validate:"required"`
	LastName          string  `json:"LastName" binding:"required,max=50" validate:"required"`
	StreetNo          string  `json:"StreetNo" binding:"required,max=50" validate:"required"`
	StreetName        string  `json:"StreetName" binding:"required,max=50" validate:"required"`
	City              string  `json:"City" binding:"required,max=50" validate:"required"`
	State             string  `json:"State" binding:"required,max=50" validate:"required"`
	PhoneNo           string  `json:"PhoneNo" binding:"required,max=50" validate:"required"`
	Zip               string  `json:"Zip" binding:"required,max=50" validate:"required"`
	Country           Country `json:"Country" binding:"required,max=50" validate:"required"`
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
