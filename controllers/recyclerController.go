package controllers

import (
	"fmt"
	"gsdc/letsfix/models"
	"gsdc/letsfix/service"
	"strconv"

	//"net/http"

	"github.com/gin-gonic/gin"
)

type RecyclerController interface {
	FindAllRecyclers() []models.Recycler
	SaveRecycler(ctx *gin.Context) error
	//Update(ctx *gin.Context) error
	//Delete(ctx *gin.Context) error
	FindByRecyclerId(ctx *gin.Context) models.Recycler
	
}

type recyclerController struct {
	service service.RecyclerService
}

func NewRecyclerController(service service.RecyclerService) RecyclerController {
	return &recyclerController {
		service: service,
	}
}

func (c *recyclerController) FindAllRecyclers() []models.Recycler {
	return c.service.FindAllRecyclers() 
}

func (c *recyclerController) SaveRecycler(ctx *gin.Context) error {
	var r models.Recycler
	err := ctx.ShouldBindJSON(&r)
	if err != nil {
		return err
	}
	c.service.SaveRecycler(r)
	return nil
}

func (c *recyclerController) FindByRecyclerId(ctx *gin.Context) models.Recycler {
	p := ctx.Param("id")
	id, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		fmt.Println(err.Error())
	}
	recycler_id := uint(id)
	return c.service.FindByRecyclerId(recycler_id)
}