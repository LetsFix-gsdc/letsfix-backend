package models

type Repair_Type struct {
	ID         uint     `gorm:"auto_increment;primary_key" json:"id,omitempty"`
	Repairer   Repairer `json:"repairer,omitempty"`
	Type       Type     `json:"type,omitempty"`
	Brand      Brand    `json:"brand,omitempty"`
	Component  string   `json:"component,omitempty"`
	MinPrice   float32  `json:"min_price,omitempty"`
	MaxPrice   float32  `json:"max_price,omitempty"`
	IsOfficial bool     `json:"is_official,omitempty"`
	IsLicensed bool     `json:"is_licensed,omitempty"`
}
