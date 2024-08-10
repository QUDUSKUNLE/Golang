package domain

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Contact struct {
	PhoneNumbers []string `json:"PhoneNumbers" validate:"required,dive,required"`
	WhatsApps    string   `json:"WhatsApp" binding:"required,max=50" validate:"required"`
	Twitter      string   `json:"Twitter" binding:"required,max=50" validate:"required"`
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
