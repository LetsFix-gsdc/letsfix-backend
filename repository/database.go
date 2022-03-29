package repository

import "gorm.io/gorm"


type database struct {
	connection *gorm.DB
}

