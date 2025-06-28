package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID      uint `gorm:"primarykey; autoIncrement"`
	Name    string
	Age     int
	Address string
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading the .env file")
	}

	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	pass := os.Getenv("PASS")
	dbname := os.Getenv("DBNAME")
	port := os.Getenv("PORT")
	sslmode := os.Getenv("SSLMODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, pass, dbname, port, sslmode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB Connection failed", err.Error())
	}
	db.AutoMigrate(&User{})
	// userdata := User{Name: "Goutham", Age: 24, Address: "Rajiv Gruha Kalpa"}
	// // create a record in db

	// if err = db.Create(&userdata).Error; err != nil {
	// 	log.Fatal("Creating database failed", err)

	// }

	// // Create a multiple records in the database;

	// userdatas := []*User{
	// 	{
	// 		Name: "Laptop", Age: 12, Address: "Pragathi nagar",
	// 	}, {
	// 		Name: "CHair", Age: 10, Address: "nizampet",
	// 	},
	// }

	// db.Create(userdatas)

	//retrieve the record with specific field

	// var userfetched User
	// db.Where(&User{Name: "Laptop", Age: 12}).First(&userfetched)

	// fmt.Printf("%+v", userfetched)

	// Get all records.
	// usersdata := []User{}
	// db.Find(&usersdata)
	// fmt.Println(usersdata)

	// Update Records
	// Update Single column;
	// db.Model(&User{}).Where("ID=?", 4).Update("Name", "Screen")

	//Update  multiple columns

	// db.Model(&User{}).Where("ID=?", 4).Updates(User{Name: "Keyboard", Age: 45})

	// Delete record;

	db.Unscoped().Delete(&User{}, 4)
}
