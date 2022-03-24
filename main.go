package main

import (
	"gsdc/letsfix/controllers"
	"gsdc/letsfix/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/users", controllers.FindUsers)
	r.POST("/users", controllers.CreateUser)

	r.Run()
}
