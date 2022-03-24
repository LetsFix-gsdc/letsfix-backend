package models

import "time"

type Ownership struct {
	ID              uint `gorm:"auto_increment;primary_key"`
	Owner           User
	Device          Device
	SerialNumber    string    `json:"serial_number"`
	PurchaseDate    time.Time `json:"purchase_date"`
	ReceiptLocation string    `json:"receipt_location"`
}
