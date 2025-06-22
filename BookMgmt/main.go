package main

import (
	"log"
	"os"

	"github.com/Goutham/BookMgmt/routes"
	"github.com/Goutham/BookMgmt/storage"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading environment variables")
	}
	router := gin.New()

	router.Use(gin.Logger())

	config := storage.Config{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		UserName: os.Getenv("USER"),
		Password: os.Getenv("PASS"),
		DBName:   os.Getenv("DB"),
		SSlMode:  os.Getenv("SSL"),
	}

	// fmt.Println(config)
	db, err := storage.NewConnection(&config)
	if err != nil {
		log.Fatal("database connection failed")
	}
	DB := storage.Repository{
		DB: db,
	}
	err = storage.MigrateBooks(db)

	if err != nil {
		log.Fatal("Error while creating a Table in a database in postgres", err.Error())
	}

	routes.SetUproutes(router, &DB)

	router.Run(":8090")

}
