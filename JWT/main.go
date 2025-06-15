package main

import (
	"fmt"
	"os"

	"github.com/Goutham/Gin/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	fmt.Println(port)
	router := gin.New()

	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.Run(":" + port)

}
