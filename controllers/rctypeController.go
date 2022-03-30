package controllers

import (
	"fmt"
	"gsdc/letsfix/models"
	"gsdc/letsfix/service"
	"strconv"

	//"net/http"

	"github.com/gin-gonic/gin"
)

type RctypeController interface {
	FindAllRctypes() []models.Recycler_Type
	SaveRctype(ctx *gin.Context) error
	//Update(ctx *gin.Context) error
	//Delete(ctx *gin.Context) error
	FindRctypeById(ctx *gin.Context) models.Recycler_Type
	FindRctypeByType(ctx *gin.Context) []models.Recycler_Type
	
}

type rctypeController struct {
	service service.RctypeService
}

func NewRctypeController(service service.RctypeService) RctypeController {
	return &rctypeController {
		service: service,
	}
}

func (c *rctypeController) FindAllRctypes() []models.Recycler_Type {
	return c.service.FindAllRctypes() 
}

func (c *rctypeController) SaveRctype(ctx *gin.Context) error {
	var r models.Recycler_Type
	err := ctx.ShouldBindJSON(&r)
	if err != nil {
		return err
	}
	c.service.SaveRctype(r)
	return nil
}

func (c *rctypeController) FindRctypeById(ctx *gin.Context) models.Recycler_Type {
	p := ctx.Param("id")
	id, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		fmt.Println(err.Error())
	}
	r_id := uint(id)
	return c.service.FindRctypeById(r_id)
}

func (c *rctypeController) FindRctypeByType(ctx *gin.Context) []models.Recycler_Type {
	name := ctx.Param("type_name")
	return c.service.FindRctypeByType(name)
}