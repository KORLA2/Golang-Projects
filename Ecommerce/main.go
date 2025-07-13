package main

import (
	"myapp/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	router := gin.New()

	router.Use(gin.Logger())
	routes.NonAuthRoutes(router, DB)
	routes.AuthRoutes(router)

	godotenv.Load(".env")
	port := os.Getenv("PORT")

	router.Run(":" + port)

}
