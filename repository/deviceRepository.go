package repository

import (
	"gsdc/letsfix/models"
	//"fmt"
	//"gsdc/letsfix/util"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	//"gorm.io/driver/postgres"
	//"gorm.io/gorm"
)

type DeviceRepository interface {
	SaveDevice(models.Device)
	UpdateDevice(models.Device)
	DeleteDevice(models.Device)
	FindAllDevices() []models.Device
	//FindDeviceByUserId()
	FindDeviceById(uint) models.Device
	//CloseDB()
}

func NewDeviceRepository() DeviceRepository {
	/*
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

		db.AutoMigrate(&models.User{}, &models.Ownership{},
		&models.Device{}, &models.Type{},
		&models.Brand{}, &models.Recycler_Type{},
		&models.Repair_Type{}, &models.Recycler{},
		&models.Repairer{})
		//db.AutoMigrate(&models.User{})

		return &database {
			connection: db,
		}
	*/
	//DB = db

	return &database{
		connection: DB,
	}
}

/*
func (db *database) CloseDB() {
	err := db.connection.
	if err != nil {
		panic("Failed to close database")
	}
}
*/

func (db *database) SaveDevice(device models.Device) {
	DB.Create(&device)
}

func (db *database) UpdateDevice(device models.Device) {
	DB.Save(&device)
}

func (db *database) DeleteDevice(device models.Device) {
	DB.Delete(&device)
}

func (db *database) FindAllDevices() []models.Device {
	var devices []models.Device
	//db.connection.Find(&device)

	//for tables with foreign keys,
	DB.Set("gorm:auto_preload", true).Find(&devices)

	return devices
}

func (db *database) FindDeviceById(device_id uint) models.Device {
	var device models.Device
	DB.Set("gorm:auto_preload", true).Find(&device, "id=?", device_id)
	return device
}
