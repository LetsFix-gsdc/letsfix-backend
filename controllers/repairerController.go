package controllers

import (
	"fmt"
	"gsdc/letsfix/models"
	"gsdc/letsfix/service"
	"strconv"

	//"net/http"

	"github.com/gin-gonic/gin"
)

type RepairerController interface {
	FindAllRepairers() []models.Repairer
	SaveRepairer(ctx *gin.Context) error
	//Update(ctx *gin.Context) error
	//Delete(ctx *gin.Context) error
	FindByRepairerId(ctx *gin.Context) models.Repairer
	
}

type repairerController struct {
	service service.RepairerService
}

func NewRepairerController(service service.RepairerService) RepairerController {
	return &repairerController {
		service: service,
	}
}

func (c *repairerController) FindAllRepairers() []models.Repairer {
	return c.service.FindAllRepairers() 
}

func (c *repairerController) SaveRepairer(ctx *gin.Context) error {
	var r models.Repairer
	err := ctx.ShouldBindJSON(&r)
	if err != nil {
		return err
	}
	c.service.SaveRepairer(r)
	return nil
}

func (c *repairerController) FindByRepairerId(ctx *gin.Context) models.Repairer {
	p := ctx.Param("id")
	id, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		fmt.Println(err.Error())
	}
	repairer_id := uint(id)
	return c.service.FindByRepairerId(repairer_id)
}