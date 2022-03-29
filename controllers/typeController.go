package controllers

import (
	"fmt"
	"gsdc/letsfix/models"
	"gsdc/letsfix/service"
	"strconv"

	//"net/http"

	"github.com/gin-gonic/gin"
)

type TypeController interface {
	FindAllTypes() []models.Type
	SaveType(ctx *gin.Context) error
	//Update(ctx *gin.Context) error
	//Delete(ctx *gin.Context) error
	FindByTypeId(ctx *gin.Context) models.Type
	
}

type typeController struct {
	service service.TypeService
}

func NewTypeController(service service.TypeService) TypeController{
	return &typeController {
		service: service,
	}
}

func (c *typeController) FindAllTypes() []models.Type {
	return c.service.FindAllTypes() 
}

func (c *typeController) SaveType(ctx *gin.Context) error {
	var t models.Type
	err := ctx.ShouldBindJSON(&t)
	if err != nil {
		return err
	}
	c.service.SaveType(t)
	return nil
}

func (c *typeController) FindByTypeId(ctx *gin.Context) models.Type {
	p := ctx.Param("id")
	id, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		fmt.Println(err.Error())
	}
	type_id := uint(id)
	return c.service.FindByTypeId(type_id)
}