package models

type Recycler_Type struct {
	ID       uint `gorm:"auto_increment;primary_key"`
	Recycler Recycler
	Type     Type
}
