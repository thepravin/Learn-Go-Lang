package main

import (
	"ginLearning/05_Auth/controllers"
	"ginLearning/05_Auth/db"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db := db.InitDB()

	if db == nil {
		log.Fatal("DB not connected.....")
	}

	authController := controllers.InitAuthController()
	authController.InitAuthControllerRoutes(router)

	router.Run(":7757")
}
