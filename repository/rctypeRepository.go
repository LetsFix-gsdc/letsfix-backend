package repository

import "gsdc/letsfix/models"

type RctypeRepository interface {
	SaveRctype(models.Recycler_Type)
	//UpdateOwnership(models.Ownership)
	//DeleteOwnership(models.Ownership)
	FindAllRctypes() []models.Recycler_Type
	//FindDeviceByUserId()
	FindRctypeById(uint) models.Recycler_Type
	FindRctypeByType(string) []models.Recycler_Type
	//CloseDB()
}


func NewRctypeRepository() RctypeRepository {
	return &database{
		connection: DB,
	}
}

func (db *database) SaveRctype(r models.Recycler_Type) {
	DB.Create(&r)
}

func (db *database) FindAllRctypes() []models.Recycler_Type {

	var r []models.Recycler_Type
	DB.Find(&r)
	/*
		for tables with foreign keys,
		d.connection.Set("gorm:auto_preload", true).Find(&users)
	*/
	return r
}

func (db *database) FindRctypeById(r_id uint) models.Recycler_Type {
	var r models.Recycler_Type
	DB.Set("gorm:auto_preload", true).Find(&r, "id=?", r_id)
	return r
}

func (db *database) FindRctypeByType(t string) []models.Recycler_Type {
	var r []models.Recycler_Type
	var rct []models.Recycler_Type
	DB.Set("gorm:auto_preload", true).Find(&r)
	for i := 0; i < len(r); i++ {
		if name := r[i].Type.Name; name == t {
			rct = append(rct, r[i])
		}
	}
 
 	return rct
}

