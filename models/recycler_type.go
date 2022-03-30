package models

type Recycler_Type struct {
	ID       uint     `gorm:"auto_increment;primary_key" json:"id,omitempty"`
	Recycler Recycler `gorm:"foreignkey:RecyclerID" json:"recycler,omitempty"`
	Type     Type     `gorm:"foreignkey:TypeID" json:"type,omitempty"`
	RecyclerID uint `json:"-"`
	TypeID uint `json:"-"`
}
