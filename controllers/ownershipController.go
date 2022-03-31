package controllers

import (
	//"fmt"
	"gsdc/letsfix/models"
	"gsdc/letsfix/service"
	//"strconv"

	//"net/http"

	"github.com/gin-gonic/gin"
)

type OwnershipController interface {
	FindAllOwnerships() []models.Ownership
	SaveOwnership(ctx *gin.Context) error
	//UpdateDevice(ctx *gin.Context) error
	//DeleteDevice(ctx *gin.Context) error
	FindOwnershipByUserId(ctx *gin.Context) []models.Ownership
	FindDevicesByUserId(ctx *gin.Context) []models.Ownership
}

type ownershipController struct {
	service service.OwnershipService
}

func NewOwnershipController(service service.OwnershipService) OwnershipController {
	return &ownershipController{
		service: service,
	}
}

func (c *ownershipController) FindAllOwnerships() []models.Ownership {
	return c.service.FindAllOwnerships()
}

func (c *ownershipController) SaveOwnership(ctx *gin.Context) error {
	var ownership models.Ownership
	err := ctx.ShouldBindJSON(&ownership)
	if err != nil {
		return err
	}
	c.service.SaveOwnership(ownership)
	return nil
}

func (c *ownershipController) FindOwnershipByUserId(ctx *gin.Context) []models.Ownership {
	user_id := ctx.Param("id")
	return c.service.FindOwnershipByUserId(user_id)
}

func (c *ownershipController) FindDevicesByUserId(ctx *gin.Context) []models.Ownership {
	user_id := ctx.Param("id")
	return c.service.FindDevicesByUserId(user_id)
}
