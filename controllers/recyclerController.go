package controllers

import (
	"fmt"
	"gsdc/letsfix/models"
	"gsdc/letsfix/service"
	"math"
	"sort"
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
	FindByLocation(ctx *gin.Context) []models.Recycler
}

type recyclerController struct {
	service service.RecyclerService
}

func NewRecyclerController(service service.RecyclerService) RecyclerController {
	return &recyclerController{
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

type RecyclerDistance struct {
	Recycler models.Recycler
	Distance float64
}

func CalculateDistance(x1 float64, x2 float64, y1 float64, y2 float64) float64 {
	xDelt := math.Abs(x1 - x2)
	yDelt := math.Abs(y1 - y2)
	distance := math.Sqrt(math.Pow(xDelt, 2) + math.Pow(yDelt, 2))
	return distance
}

func (c *recyclerController) FindByLocation(ctx *gin.Context) []models.Recycler {
	const RESULT_LIMIT int = 25
	lat, _ := strconv.ParseFloat(ctx.Param("lat"), 64)
	long, _ := strconv.ParseFloat(ctx.Param("long"), 64)

	var distances []RecyclerDistance
	recyclers := c.service.FindAllRecyclers()

	for _, r := range recyclers {
		rd := RecyclerDistance{
			Recycler: r,
			Distance: CalculateDistance(r.Latitude, lat, r.Longtitude, long),
		}
		distances = append(distances, rd)
	}

	sort.Slice(distances[:], func(i, j int) bool {
		return distances[i].Distance < distances[j].Distance
	})

	var res []models.Recycler

	for _, r := range distances {
		res = append(res, r.Recycler)
	}

	return res[:RESULT_LIMIT]
}
