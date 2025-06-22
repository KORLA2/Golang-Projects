package storage

import (
	"fmt"

	"github.com/Goutham/BookMgmt/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

type Config struct {
	Host     string
	Port     string
	UserName string
	Password string
	DBName   string
	SSlMode  string
}

func MigrateBooks(db *gorm.DB) error {

	err := db.AutoMigrate(&model.Book{})

	return err
}

func NewConnection(config *Config) (*gorm.DB, error) {
	fmt.Println(*config)
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", config.Host, config.Port, config.UserName, config.Password, config.DBName, config.SSlMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil

}
