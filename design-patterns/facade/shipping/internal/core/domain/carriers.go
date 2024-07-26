package domain

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Carrier struct {
	gorm.Model
	ID        			uuid.UUID  `json:"ID" gorm:"uuid;primaryKey"`
	CreatedAt 			time.Time  `json:"CreatedAt"`
	UpdatedAt 			*time.Time `json:"UpdatedAt,omitempty"`

	CompanyName 		string   		`json:"CompanyName"`
	CompanyAddress 	Address  		`gorm:"embedded" json:"CompanyAddress"`
	Contact       	Contact  		`gorm:"embedded" json:"Address"`
	UserID    			uuid.UUID 	`json:"CarrierID"`
	User      			*User
	
	PickUps   			[]PickUp   `json:"PickUps" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type CarrierDto struct {
	CompanyName     string `json:"CompanyName" binding:"required" validate:"required,gte=6,lte=1000"`
	CompanyAddress  Address `json:"CompanyAddress" binding:"required" validate:"required"`
	Contact         Contact  `json:"Contact" binding:"required" validate:"required"`
	UserID          uuid.UUID
}

func (carrier *Carrier) BeforeCreate(scope *gorm.DB) error {
	carrier.ID = uuid.New()
	return nil
}


func (carr *Carrier) BuildNewCarrier(carrierDto CarrierDto) (*Carrier, error) {
	return &Carrier{
		UserID: carrierDto.UserID,
	}, nil
}
