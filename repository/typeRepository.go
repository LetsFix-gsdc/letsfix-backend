package repository

import "gsdc/letsfix/models"

type TypeRepository interface {
	SaveType(models.Type)
	//UpdateOwnership(models.Ownership)
	//DeleteOwnership(models.Ownership)
	FindAllTypes() []models.Type
	//FindDeviceByUserId()
	FindByTypeId(uint) models.Type
	//CloseDB()
}


func NewTypeRepository() TypeRepository {
	return &database{
		connection: DB,
	}
}

func (db *database) SaveType(t models.Type) {
	DB.Create(&t)
}

func (db *database) FindAllTypes() []models.Type {

	var types []models.Type
	DB.Find(&types)
	/*
		for tables with foreign keys,
		d.connection.Set("gorm:auto_preload", true).Find(&users)
	*/
	return types
}

func (db *database) FindByTypeId(type_id uint) models.Type {
	var t models.Type
	DB.Find(&t, "id=?", type_id)
	return t
}