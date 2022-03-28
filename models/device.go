package models

import "time"

type Device struct {
	ID               uint          `gorm:"auto_increment;primary_key" json:"id,omitempty"`
	Name             string        `json:"name,omitempty"`
	Type             Type          `gorm:"foreignkey:TypeName" json:"type,omitempty"`
	Brand            Brand         `gorm:"foreignkey:BrandName" json:"brand,omitempty"`
	WarrantyDuration time.Duration `json:"warranty_duration,omitempty"`
	TypeName string `json:"-"`
	BrandName string `json:"-"`
}
