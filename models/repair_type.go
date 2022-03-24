package models

type Repair_Type struct {
	ID         uint `gorm:"auto_increment;primary_key"`
	Repairer   Repairer
	Type       Type
	Brand      Brand
	Component  string
	MinPrice   float32
	MaxPrice   float32
	IsOfficial bool
	IsLicensed bool
}
