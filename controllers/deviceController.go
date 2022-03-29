package controllers

import (
	"fmt"
	"gsdc/letsfix/models"
	"gsdc/letsfix/service"
	"strconv"

	//"net/http"

	"github.com/gin-gonic/gin"
)

type DeviceController interface {
	FindAllDevices() []models.Device
	SaveDevice(ctx *gin.Context) error
	UpdateDevice(ctx *gin.Context) error
	DeleteDevice(ctx *gin.Context) error
	FindDevice(ctx *gin.Context) error
	
}

type deviceController struct {
	service service.DeviceService
}

func NewDeviceController(service service.DeviceService) DeviceController{
	return &deviceController {
		service: service,
	}
}

func (c *deviceController) FindAllDevices() []models.Device {
	return c.service.FindAllDevices() 
}

func (c *deviceController) SaveDevice(ctx *gin.Context) error {
	var device models.Device
	err := ctx.ShouldBindJSON(&device)
	if err != nil {
		return err
	}
	c.service.SaveDevice(device)
	return nil
}

func (c *deviceController) UpdateDevice(ctx *gin.Context) error {
	var device models.Device
	err := ctx.ShouldBindJSON(&device)
	if err != nil {
		return err
	}
	p := ctx.Param("id")
	device_id, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		fmt.Println(err.Error())
	}
	device.ID = uint(device_id)
	c.service.UpdateDevice(device)
	return nil
}

func (c *deviceController) DeleteDevice(ctx *gin.Context) error {
	var device models.Device
	p := ctx.Param("id")
	device_id, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		fmt.Println(err.Error())
	}
	device.ID = uint(device_id)
	c.service.DeleteDevice(device)
	return nil

}

func (c *deviceController) FindDevice(ctx *gin.Context) error {
	p := ctx.Param("id")
	id, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		fmt.Println(err.Error())
	}
	device_id := uint(id)
	c.service.FindDevice(device_id)
	return nil
}


