package controllers

import (
	"fmt"
	"gsdc/letsfix/models"
	"gsdc/letsfix/service"

	//"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	FindAll() []models.User
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
	FindByUserId(ctx *gin.Context) models.User
	
}

type userController struct {
	service service.UserService
}

func New(service service.UserService) UserController{
	return &userController {
		service: service,
	}
}

func (c *userController) FindAll() []models.User {
	return c.service.FindAll() 
}

func (c *userController) Save(ctx *gin.Context) error {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return err
	}
	c.service.Save(user)
	return nil
}

func (c *userController) Update(ctx *gin.Context) error {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return err
	}
	user_id := ctx.Param("id")
	user.ID = user_id
	c.service.Update(user)
	return nil
}

func (c *userController) Delete(ctx *gin.Context) error {
	var user models.User
	user_id := ctx.Param("id")
	user.ID = user_id
	c.service.Delete(user)
	return nil

}

func (c *userController) FindByUserId(ctx *gin.Context) models.User {
	user_id := ctx.Param("id")
	fmt.Println(user_id)
	return c.service.FindByUserId(user_id)
}


/*
func FindUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func CreateUser(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Create(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}
*/