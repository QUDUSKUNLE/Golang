package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type (
	Packaging struct {
		gorm.Model
		ID        uuid.UUID      `gorm:"primaryKey;->;<-:create" json:"id"`
		CreatedAt time.Time      `json:"created_at"`
		UpdatedAt time.Time      `json:"updated_at"`
		DeletedAt gorm.DeletedAt `json:"deleted_at"`

		UserID              uuid.UUID `json:"-"`
		User                *User     `json:"-"`
		TerminalPackagingID string    `json:"terminal_packaging_id"`
	}
	PackagingDto struct {
		UserID              uuid.UUID
		TerminalPackagingID string `json:"terminal_packaging_id"`
	}
)

func (packaging *Packaging) BeforeCreate(scope *gorm.DB) error {
	packaging.ID = uuid.New()
	return nil
}

func (packaging *Packaging) BuildNewPackaging(packageDto PackagingDto) *PackagingDto {
	return &PackagingDto{
		UserID:              packageDto.UserID,
		TerminalPackagingID: packageDto.TerminalPackagingID,
	}
}
