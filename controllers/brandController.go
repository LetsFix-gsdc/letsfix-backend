package controllers

import (
	"fmt"
	"gsdc/letsfix/models"
	"gsdc/letsfix/service"
	"strconv"

	//"net/http"

	"github.com/gin-gonic/gin"
)

type BrandController interface {
	FindAllBrands() []models.Brand
	SaveBrand(ctx *gin.Context) error
	//Update(ctx *gin.Context) error
	//Delete(ctx *gin.Context) error
	FindByBrandId(ctx *gin.Context) models.Brand
	
}

type brandController struct {
	service service.BrandService
}

func NewBrandController(service service.BrandService) BrandController{
	return &brandController {
		service: service,
	}
}

func (c *brandController) FindAllBrands() []models.Brand {
	return c.service.FindAllBrands() 
}

func (c *brandController) SaveBrand(ctx *gin.Context) error {
	var b models.Brand
	err := ctx.ShouldBindJSON(&b)
	if err != nil {
		return err
	}
	c.service.SaveBrand(b)
	return nil
}

func (c *brandController) FindByBrandId(ctx *gin.Context) models.Brand {
	p := ctx.Param("id")
	id, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		fmt.Println(err.Error())
	}
	brand_id := uint(id)
	return c.service.FindByBrandId(brand_id)
}