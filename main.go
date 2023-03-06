package main

import (
	"github.com/ahmadnurrizal/go-gin-jwt-authentication/controllers"
	"github.com/ahmadnurrizal/go-gin-jwt-authentication/models"
	"github.com/gin-gonic/gin"
)

func main() {

	models.ConnectDataBase()

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	r.Run("localhost:8080")

}
