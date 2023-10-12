package controllers

import (
	"api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var newUser models.User
	db := models.ConnectDB()

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := db.Create(&newUser); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": newUser})
}

func DeleteUser(c *gin.Context) {
	db := models.ConnectDB()
	var user models.User

	if err := db.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	db.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"data": "User deleted!"})
}
