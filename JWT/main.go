package main

import (
	"log"

	"github.com/Goutham/Gin/database"
	"github.com/Goutham/Gin/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	router := gin.New()

	router.Use(gin.Logger())

	db, err := database.NewConnection()

	if err != nil {
		log.Fatal("Unable to establish connection with postgres", err.Error())
	}

	routes.AuthRoutes(router, db)
	routes.UserRoutes(router, db)

	router.Run(":8090")

}
