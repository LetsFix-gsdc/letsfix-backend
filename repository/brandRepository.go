package repository

import "gsdc/letsfix/models"

type BrandRepository interface {
	SaveBrand(models.Brand)
	//UpdateOwnership(models.Ownership)
	//DeleteOwnership(models.Ownership)
	FindAllBrands() []models.Brand
	//FindDeviceByUserId()
	FindByBrandId(uint) models.Brand
	//CloseDB()
}


func NewBrandRepository() BrandRepository {
	return &database{
		connection: DB,
	}
}

func (db *database) SaveBrand(b models.Brand) {
	DB.Create(&b)
}

func (db *database) FindAllBrands() []models.Brand {

	var brands []models.Brand
	DB.Find(&brands)
	/*
		for tables with foreign keys,
		d.connection.Set("gorm:auto_preload", true).Find(&users)
	*/
	return brands
}

func (db *database) FindByBrandId(brand_id uint) models.Brand {
	var b models.Brand
	DB.Find(&b, "id=?", brand_id)
	return b
}