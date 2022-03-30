package controllers

import (
	"fmt"
	"gsdc/letsfix/models"
	"gsdc/letsfix/service"
	"strconv"

	//"net/http"

	"github.com/gin-gonic/gin"
)

type RptypeController interface {
	FindAllRptypes() []models.Repair_Type
	SaveRptype(ctx *gin.Context) error
	FindRptypeById(ctx *gin.Context) models.Repair_Type
	FindRptypeByRepairer(ctx *gin.Context) []models.Repair_Type
	FindRptypeByType(ctx *gin.Context) []models.Repair_Type
	FindRptypeByBrand(ctx *gin.Context) []models.Repair_Type
	FindRptypeByComponent(ctx *gin.Context) []models.Repair_Type
}

type rptypeController struct {
	service service.RptypeService
}

func NewRptypeController(service service.RptypeService) RptypeController {
	return &rptypeController {
		service: service,
	}
}

func (c *rptypeController) FindAllRptypes() []models.Repair_Type {
	return c.service.FindAllRptypes() 
}

func (c *rptypeController) SaveRptype(ctx *gin.Context) error {
	var r models.Repair_Type
	err := ctx.ShouldBindJSON(&r)
	if err != nil {
		return err
	}
	c.service.SaveRptype(r)
	return nil
}

func (c *rptypeController) FindRptypeById(ctx *gin.Context) models.Repair_Type {
	p := ctx.Param("id")
	id, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		fmt.Println(err.Error())
	}
	r_id := uint(id)
	return c.service.FindRptypeById(r_id)
}

func (c *rptypeController) FindRptypeByRepairer(ctx *gin.Context) []models.Repair_Type {
	name := ctx.Param("repairer")
	return c.service.FindRptypeByRepairer(name)
}

func (c *rptypeController) FindRptypeByType(ctx *gin.Context) []models.Repair_Type {
	name := ctx.Param("type_name")
	return c.service.FindRptypeByType(name)
}

func (c *rptypeController) FindRptypeByBrand(ctx *gin.Context) []models.Repair_Type {
	name := ctx.Param("brand_name")
	return c.service.FindRptypeByBrand(name)
}

func (c *rptypeController) FindRptypeByComponent(ctx *gin.Context) []models.Repair_Type {
	name := ctx.Param("component")
	return c.service.FindRptypeByComponent(name)
}