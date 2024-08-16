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
		CreatedAt *time.Time      `json:"created_at"`
		UpdatedAt *time.Time      `json:"updated_at"`
		DeletedAt *gorm.DeletedAt `json:"-"`

		UserID              uuid.UUID `json:"-"`
		User                *User     `json:"-"`
		TerminalPackagingID string    `json:"terminal_packaging_id"`
	}
	PackagingDto struct {
		UserID              uuid.UUID
		PackagingID []string `json:"packaging_ids"`
	}
)

func (packaging *Packaging) BeforeCreate(scope *gorm.DB) error {
	packaging.ID = uuid.New()
	return nil
}

func (packaging *Packaging) BuildNewPackaging(packageDto PackagingDto) []*Packaging {
	pack := []*Packaging{}
	for _, terminal_id := range packageDto.PackagingID {
		pack = append(pack, &Packaging{
			UserID: packageDto.UserID,
			TerminalPackagingID: terminal_id,
		})
	}
	return pack
}
