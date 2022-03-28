package models

import "time"

type Ownership struct {
	ID              uint      `gorm:"auto_increment;primary_key" json:"id,omitempty"`
	Owner           User      `json:"owner,omitempty"`
	Device          Device    `json:"device,omitempty"`
	SerialNumber    string    `json:"serial_number,omitempty"`
	PurchaseDate    time.Time `json:"purchase_date,omitempty"`
	ReceiptLocation string    `json:"receipt_location,omitempty"`
	/*
	ID              uint `gorm:"auto_increment;primary_key" json:"id,omitempty"`
	Owner           User `gorm:"foreignkey:UserID"`
	Device          Device `gorm:"foreignkey:DeviceID"`
	SerialNumber    string    `json:"serial_number"`
	PurchaseDate    time.Time `json:"purchase_date"`
	ReceiptLocation string    `json:"receipt_location"`
	UserID string `json:"-"`
	DeviceID uint `json:"-"`
	*/
}
