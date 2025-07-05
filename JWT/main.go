package main

import (
	"log"
	"os"

	"github.com/Goutham/Gin/database"
	"github.com/Goutham/Gin/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")

	router := gin.New()

	router.Use(gin.Logger())

	db, err := database.NewConnection()

	if err != nil {
		log.Fatal("Unable to establish connection with postgres", err.Error())
	}
	routes.AuthRoutes(router,db)
	routes.UserRoutes(router,db)

	router.Run(":", port)

}
