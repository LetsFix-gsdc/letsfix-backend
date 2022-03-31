package repository

import "gsdc/letsfix/models"

type OwnershipRepository interface {
	SaveOwnership(models.Ownership)
	//UpdateOwnership(models.Ownership)
	//DeleteOwnership(models.Ownership)
	FindAllOwnerships() []models.Ownership
	//FindDeviceByUserId()
	FindOwnershipByUserId(string) []models.Ownership
	FindDevicesByUserId(string) []models.Ownership
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
		if d := ownerships[i].UserID; d == user_id {
			o = append(o, ownerships[i])
		}
	}

	return o
}

func (db *database) FindDevicesByUserId(user_id string) []models.Ownership {
	var allOwnerships []models.Ownership
	var res []models.Ownership

	DB.Set("gorm:auto_preload", true).Find(&allOwnerships)
	for i := 0; i < len(allOwnerships); i++ {
		if d := allOwnerships[i].UserID; d == user_id {
			res = append(res, allOwnerships[i])
		}
	}
	return res

}
