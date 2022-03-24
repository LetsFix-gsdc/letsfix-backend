package models

import "time"

type Device struct {
	ID               uint   `gorm:"auto_increment;primary_key"`
	Name             string `json:"name"`
	Type             Type
	Brand            Brand
	WarrantyDuration time.Duration
}
