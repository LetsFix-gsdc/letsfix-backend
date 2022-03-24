package models

type Type struct {
	ID   uint `gorm:"auto_increment;primary_key"`
	Name string
}
