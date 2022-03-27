package models

type Recycler_Type struct {
	ID       uint     `gorm:"auto_increment;primary_key" json:"id,omitempty"`
	Recycler Recycler `json:"recycler,omitempty"`
	Type     Type     `json:"type,omitempty"`
}
