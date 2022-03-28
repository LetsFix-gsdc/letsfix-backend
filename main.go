package main

import (
	"gsdc/letsfix/controllers"
	//"gsdc/letsfix/models"
	"gsdc/letsfix/repository"
	"gsdc/letsfix/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	userRepository repository.UserRepository = repository.NewUserRepository()
	userService service.UserService = service.New(userRepository)
	userController controllers.UserController = controllers.New(userService)
)

func main() {
	r := gin.Default()
	//models.ConnectDatabase()
	
	/*
	r.GET("/users", controllers.FindUsers)
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:user_id/devices", controllers.FindDevices)
	*/

	
	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, userController.FindAll())
	})

	r.POST("/users", func(ctx *gin.Context) {
		err := userController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "User Input is Valid!"})
		}
	})

	r.PUT("/users/:id", func(ctx *gin.Context) {
		err := userController.Update(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "User Update is Valid!"})
		}
	})
	
	

	r.Run()
}
