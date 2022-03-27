package models

type Type struct {
	ID        uint   `gorm:"auto_increment;primary_key" json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Flowchart string `json:"flowchart,omitempty"`
}
