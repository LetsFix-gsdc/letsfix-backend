package models

import "time"

type Device struct {
	ID               uint          `gorm:"auto_increment;primary_key" json:"id,omitempty"`
	Name             string        `json:"name,omitempty"`
	Type             Type          `gorm:"foreignkey:TypeID" json:"type,omitempty"`
	Brand            Brand         `gorm:"foreignkey:BrandID" json:"brand,omitempty"`
	WarrantyDuration time.Duration `json:"warranty_duration,omitempty"`
	TypeID uint `json:"-"`
	BrandID uint `json:"-"`
}
