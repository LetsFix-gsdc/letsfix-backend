package repository

import "gsdc/letsfix/models"

type OwnershipRepository interface {
	SaveOwnership(models.Ownership)
	//UpdateOwnership(models.Ownership)
	//DeleteOwnership(models.Ownership)
	FindAllOwnerships() []models.Ownership
	//FindDeviceByUserId()
	FindOwnershipByUserId(string) models.Ownership
	//CloseDB()
}


func NewOwnershipRepository() OwnershipRepository {
	return &database{
		connection: DB,
	}
}

func (db *database) SaveOwnership(ownership models.Ownership) {
	DB.Create(&ownership)
}


func (db *database) FindAllOwnerships() []models.Ownership {
	var ownerships []models.Ownership
	//db.connection.Find(&device)

	//for tables with foreign keys,
	DB.Set("gorm:auto_preload", true).Find(&ownerships)

	return ownerships
}

func (db *database) FindOwnershipByUserId(user_id string) models.Ownership {
	var owner models.Ownership
	DB.Set("gorm:auto_preload", true).Find(&owner, user_id)
	return owner
}