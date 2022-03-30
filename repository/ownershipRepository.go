package repository

import "gsdc/letsfix/models"

type OwnershipRepository interface {
	SaveOwnership(models.Ownership)
	//UpdateOwnership(models.Ownership)
	//DeleteOwnership(models.Ownership)
	FindAllOwnerships() []models.Ownership
	//FindDeviceByUserId()
	FindOwnershipByUserId(string) []models.Ownership
	FindDevicesByUserId(string) []models.Device
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

func (db *database) FindOwnershipByUserId(user_id string) []models.Ownership {
	var ownerships []models.Ownership
	var o []models.Ownership
	DB.Set("gorm:auto_preload", true).Find(&ownerships)
	for i := 0; i < len(ownerships); i++ {
		if d := ownerships[i].Owner.ID; d == user_id {
			o = append(o, ownerships[i])
		}
	}

	return o
}

func (db *database) FindDevicesByUserId(user_id string) []models.Device {
	var ownerships []models.Ownership
	var allDevices []models.Device

	DB.Set("gorm:auto_preload", true).Find(&ownerships)
	for i := 0; i < len(ownerships); i++ {
		if d := ownerships[i].Owner.ID; d == user_id {
			allDevices = append(allDevices, ownerships[i].Device)
		}
	}
	return allDevices
		
}
