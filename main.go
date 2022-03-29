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

	deviceRepository repository.DeviceRepository = repository.NewDeviceRepository()
	deviceService service.DeviceService = service.NewDeviceService(deviceRepository)
	deviceController controllers.DeviceController = controllers.NewDeviceController(deviceService)

	ownershipRepository repository.OwnershipRepository = repository.NewOwnershipRepository()
	ownershipService service.OwnershipService = service.NewOwnershipService(ownershipRepository)
	ownershipController controllers.OwnershipController = controllers.NewOwnershipController(ownershipService)
)

func main() {
	r := gin.Default()
	//models.ConnectDatabase()
	
	/*
	r.GET("/users", controllers.FindUsers)
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:user_id/devices", controllers.FindDevices)
	*/

	repository.DatabaseConnect()

	// Users CRUD
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
	
	r.GET("users/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, userController.FindByUserId(ctx))
	})

	r.PUT("/users/:id", func(ctx *gin.Context) {
		err := userController.Update(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "User Update is Valid!"})
		}
	})

	// Devices CRUD

	r.GET("/devices", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, deviceController.FindAllDevices())
	})

	r.GET("/devices/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, deviceController.FindDevice(ctx))
	})

	r.POST("/devices", func(ctx *gin.Context) {
		err := deviceController.SaveDevice(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "User Input is Valid!"})
		}
	})

	// Ownership CRUD
	r.GET("/ownerships", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, ownershipController.FindAllOwnerships())
	})

	r.GET("ownerships/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, ownershipController.FindOwnershipByUserId(ctx))
	})

	r.POST("/ownerships", func(ctx *gin.Context) {
		err := ownershipController.SaveOwnership(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "User Input is Valid!"})
		}
	})	

	



	
	

	r.Run()
}
