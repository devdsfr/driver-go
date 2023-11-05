package main

import (
	"api/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/users", controllers.CreateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	r.GET("/vehicles", controllers.GetAllVehicles)
	r.POST("/vehicles", controllers.CreateVehicle)
	r.PUT("/vehicles/:id", controllers.UpdateVehicle) // Rota para atualizar um veículo
	r.DELETE("/vehicles/:id", controllers.DeleteVehicle)

	r.POST("/vehicles/:id/rent", controllers.RentVehicle) // Exemplo de rota para alugar um veículo
	r.POST("/rents", controllers.CreateRent)
	r.POST("/rents/:id", controllers.DeleteRent)

	err := r.Run()
	if err != nil {
		return
	}
}
