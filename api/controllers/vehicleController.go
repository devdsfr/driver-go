package controllers

import (
	"api/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetAllVehicles(c *gin.Context) {
	db := models.ConnectDB()
	var vehicles []models.Vehicle

	db.Find(&vehicles)
	log.Println("Buscando todos os veiculos ")

	c.JSON(http.StatusOK, gin.H{"data": vehicles})
}

func CreateVehicle(c *gin.Context) {
	db := models.ConnectDB()
	var vehicle models.Vehicle

	if err := c.ShouldBindJSON(&vehicle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&vehicle)
	c.JSON(http.StatusOK, gin.H{"data": vehicle})
}

// Atualizar veículo
func UpdateVehicle(c *gin.Context) {
	db := models.ConnectDB()
	var vehicle models.Vehicle

	if err := db.First(&vehicle, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Vehicle not found!"})
		return
	}

	if err := c.ShouldBindJSON(&vehicle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Save(&vehicle)
	c.JSON(http.StatusOK, gin.H{"data": vehicle})
}

// Deletar veículo
func DeleteVehicle(c *gin.Context) {
	db := models.ConnectDB()
	var vehicle models.Vehicle

	if err := db.First(&vehicle, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Vehicle not found!"})
		return
	}

	db.Delete(&vehicle)
	c.JSON(http.StatusOK, gin.H{"data": "Vehicle deleted!"})
}

func RentVehicle(c *gin.Context) {
	db := models.ConnectDB()
	var vehicle models.Vehicle

	if err := db.First(&vehicle, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Vehicle not found!"})
		return
	}

	user := models.User{} // Assuma que você obteve o usuário de algum middleware ou sessão

	vehicle.Status = "rented"
	vehicle.UserID = &user.ID

	db.Save(&vehicle)
	c.JSON(http.StatusOK, gin.H{"message": "Vehicle rented!", "data": vehicle})
}
