package models

type Brand struct {
	ID      uint `gorm:"auto_increment;primary_key"`
	Name    string
	Email   string
	Phone   uint
	Address string
}
