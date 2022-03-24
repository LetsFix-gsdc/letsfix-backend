package models

import (
	"fmt"
	"gsdc/letsfix/util"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
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

	db.AutoMigrate(&User{})

	DB = db
}
