package repository

import "gsdc/letsfix/models"

type RptypeRepository interface {
	SaveRptype(models.Repair_Type)
	FindAllRptypes() []models.Repair_Type
	FindRptypeById(uint) models.Repair_Type
	FindRptypeByRepairer(string) []models.Repair_Type
	FindRptypeByType(string) []models.Repair_Type
	FindRptypeByBrand(string) []models.Repair_Type
	FindRptypeByComponent(string) []models.Repair_Type
	
}


func NewRptypeRepository() RptypeRepository {
	return &database{
		connection: DB,
	}
}

func (db *database) SaveRptype(r models.Repair_Type) {
	DB.Create(&r)
}

func (db *database) FindAllRptypes() []models.Repair_Type {

	var r []models.Repair_Type
	DB.Set("gorm:auto_preload", true).Find(&r)
	/*
		for tables with foreign keys,
		d.connection.Set("gorm:auto_preload", true).Find(&users)
	*/
	return r
}

func (db *database) FindRptypeById(r_id uint) models.Repair_Type {
	var r models.Repair_Type
	DB.Set("gorm:auto_preload", true).Find(&r, "id=?", r_id)
	return r
}

func (db *database) FindRptypeByRepairer(re string) []models.Repair_Type {
	var r []models.Repair_Type
	var rpt []models.Repair_Type
	DB.Set("gorm:auto_preload", true).Find(&r)
	
	for i := 0; i < len(r); i++ {
		if rep := r[i].Repairer.Name; rep == re {
			rpt = append(rpt, r[i])
		}
	}
 
 	return r
}

func (db *database) FindRptypeByType(t string) []models.Repair_Type {
	var r []models.Repair_Type
	var rpt []models.Repair_Type
	DB.Set("gorm:auto_preload", true).Find(&r)
	for i := 0; i < len(r); i++ {
		if name := r[i].Type.Name; name == t {
			rpt = append(rpt, r[i])
		}
	}
 
 	return rpt
}

func (db *database) FindRptypeByBrand(b string) []models.Repair_Type {
	var r []models.Repair_Type
	var rpt []models.Repair_Type
	DB.Set("gorm:auto_preload", true).Find(&r)
	for i := 0; i < len(r); i++ {
		if name := r[i].Brand.Name; name == b {
			rpt = append(rpt, r[i])
		}
	}
 
 	return rpt
}

func (db *database) FindRptypeByComponent(c string) []models.Repair_Type {
	var r []models.Repair_Type
	var rpt []models.Repair_Type
	DB.Set("gorm:auto_preload", true).Find(&r)
	for i := 0; i < len(r); i++ {
		if name := r[i].Component; name == c {
			rpt = append(rpt, r[i])
		}
	}
 
 	return rpt
}