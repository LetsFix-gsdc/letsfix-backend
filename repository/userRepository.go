package repository

import (
	"gsdc/letsfix/models"
	"fmt"
	"gsdc/letsfix/util"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(models.User)
	Update(models.User)
	Delete(models.User)
	FindAll() []models.User
	//CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewUserRepository() UserRepository {

	config, err := util.LoadConfig("./.")
	if err != nil {
		panic("cannot load config: " + err.Error())
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", config.DbHost, config.DbPort, config.DbUser, config.DbPassword)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}))

	if err != nil {
		panic("Failed to connect to database!")
	}
	/*
	db.AutoMigrate(&models.User{}, &models.Ownership{}, 
	&models.Device{}, &models.Type{}, 
	&models.Brand{}, &models.Recycler_Type{},
	&models.Repair_Type{}, &models.Recycler{},
	&models.Repairer{})*/
	db.AutoMigrate(&models.User{})

	return &database {
		connection: db,
	}

	//DB = db
}
/*
func (db *database) CloseDB() {
	err := db.connection.
	if err != nil {
		panic("Failed to close database")
	}
}*/

func (db *database) Save(user models.User) {
	db.connection.Create(&user)
}

func (db *database) Update(user models.User) {
	db.connection.Save(&user)
}

func (db *database) Delete(user models.User) {
	db.connection.Delete(&user)
}

func (db *database) FindAll() []models.User {
	var users []models.User
	db.connection.Find(&users)
	/*
	for tables with foreign keys,
	d.connection.Set("gorm:auto_preload", true).Find(&users)
	*/
	return users
}



