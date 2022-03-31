package models

import "time"

type Ownership struct {
	ID             uint      `gorm:"auto_increment;primary_key" json:"id,omitempty"`
	UserID         string    `json:"user_id,omitempty"`
	DeviceID       uint      `json:"device_id,omitempty"`
	DeviceName     string    `json:"device_name,omitempty"`
	DeviceBrand    string    `json:"device_brand,omitempty"`
	SerialNumber   string    `json:"serial_number,omitempty"`
	PurchaseDate   time.Time `json:"purchase_date,omitempty"`
	WarrantyNumber string    `json:"warranty_number,omitempty"`
	//Owner           User      `gorm:"foreignkey:UserID" json:"owner,omitempty"`
	//Device          Device    `gorm:"foreignkey:DeviceID" json:"device,omitempty"`
	//ReceiptLocation string    `json:"receipt_location,omitempty"`
}
