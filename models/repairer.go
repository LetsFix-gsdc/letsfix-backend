package models

type Repairer struct {
	ID         uint    `gorm:"auto_increment;primary_key" json:"id,omitempty"`
	Name       string  `json:"name,omitempty"`
	Email      string  `json:"email,omitempty"`
	Phone      uint    `json:"phone,omitempty"`
	Address    string  `json:"address,omitempty"`
	Latitude   float64 `json:"latitude,omitempty"`
	Longtitude float64 `json:"longtitude,omitempty"`
}
