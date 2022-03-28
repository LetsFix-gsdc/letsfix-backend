package main

import (
	"gsdc/letsfix/controllers"
	"gsdc/letsfix/models"
	"gsdc/letsfix/service"

	"github.com/gin-gonic/gin"
)

var (
	userService service.UserService = service.New()
	userController controllers.UserController = controllers.New(userService)
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()
	
	/*
	r.GET("/users", controllers.FindUsers)
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:user_id/devices", controllers.FindDevices)
	*/

	
	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, userController.FindAll())
	})

	r.POST("/users", func(ctx *gin.Context) {
		ctx.JSON(200, userController.Save(ctx))
	})
	
	

	r.Run()
}
