package main

import (
	// "login-signup-api/controllers"
	"login-signup-api/config"
	"login-signup-api/routes"

	"github.com/gin-gonic/gin"

	"os"
)

func main() {

	router := gin.New()
	config.Connect()
	// models.Connect()
	// public := router.Group("/api")

	// public.POST("/register", controllers.Register)
	routes.UserRoute(router)
	routes.DashboardRoute(router)
	routes.ManageAccess(router)
	router.Run(":" + os.Getenv("PORT"))

}