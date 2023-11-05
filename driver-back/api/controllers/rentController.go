package controllers

import (
	"api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRent(c *gin.Context) {
	db := models.ConnectDB()
	var rent models.Rent

	if err := c.ShouldBindJSON(&rent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&rent)
	c.JSON(http.StatusOK, gin.H{"data": rent})
}

func DeleteRent(c *gin.Context) {
	db := models.ConnectDB()
	var rent models.Rent

	if err := db.First(&rent, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Rent not found!"})
		return
	}

	db.Delete(&rent)
	c.JSON(http.StatusOK, gin.H{"data": "Rent deleted!"})
}
