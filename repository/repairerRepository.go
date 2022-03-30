package repository

import "gsdc/letsfix/models"

type RepairerRepository interface {
	SaveRepairer(models.Repairer)
	//UpdateOwnership(models.Ownership)
	//DeleteOwnership(models.Ownership)
	FindAllRepairers() []models.Repairer
	//FindDeviceByUserId()
	FindByRepairerId(uint) models.Repairer
	//CloseDB()
}


func NewRepairerRepository() RepairerRepository {
	return &database{
		connection: DB,
	}
}

func (db *database) SaveRepairer(r models.Repairer) {
	DB.Create(&r)
}

func (db *database) FindAllRepairers() []models.Repairer {

	var repairers []models.Repairer
	DB.Find(&repairers)
	/*
		for tables with foreign keys,
		d.connection.Set("gorm:auto_preload", true).Find(&users)
	*/
	return repairers
}

func (db *database) FindByRepairerId(r_id uint) models.Repairer {
	var r models.Repairer
	DB.Find(&r, "id=?", r_id)
	return r
}