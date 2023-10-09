package main

import (
	"api/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/users", controllers.CreateUser)

	r.Run() // Ouve na porta padrão 8080
}
