package repository

import "gsdc/letsfix/models"

type RecyclerRepository interface {
	SaveRecycler(models.Recycler)
	//UpdateOwnership(models.Ownership)
	//DeleteOwnership(models.Ownership)
	FindAllRecyclers() []models.Recycler
	//FindDeviceByUserId()
	FindByRecyclerId(uint) models.Recycler
	//CloseDB()
}


func NewRecyclerRepository() RecyclerRepository {
	return &database{
		connection: DB,
	}
}

func (db *database) SaveRecycler(r models.Recycler) {
	DB.Create(&r)
}

func (db *database) FindAllRecyclers() []models.Recycler {

	var recyclers []models.Recycler
	DB.Find(&recyclers)
	/*
		for tables with foreign keys,
		d.connection.Set("gorm:auto_preload", true).Find(&users)
	*/
	return recyclers
}

func (db *database) FindByRecyclerId(r_id uint) models.Recycler {
	var r models.Recycler
	DB.Find(&r, "id=?", r_id)
	return r
}