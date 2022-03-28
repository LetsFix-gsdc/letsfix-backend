package controllers

import (
	"gsdc/letsfix/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// welcome?firstname=Jane&lastname=Doe
// /users/:user_id/devices
func FindDevices(c *gin.Context) {
	var devices []models.Device
	var users []models.Ownership

	user_id := c.Param("user_id")

	models.DB.Find(&users)
	for i := 0; i < len(users); i++ {
		if id := users[i].Owner.ID; id == user_id {
			devices = append(devices, users[i].Device)
		} 
	}
	c.JSON(http.StatusOK, gin.H{"data": devices})

}

// /users/:user_id/devices
func CreateDevice(c *gin.Context) {
	var input models.Device

	
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