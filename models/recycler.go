package models

type Recycler struct {
	ID      uint   `gorm:"auto_increment;primary_key" json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Email   string `json:"email,omitempty"`
	Phone   uint   `json:"phone,omitempty"`
	Address string `json:"address,omitempty"`
}
