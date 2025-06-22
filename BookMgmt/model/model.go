package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ID        int64   `gorm:"primary key; autoIncrement" json:"id"`
	Title     string  `json:"title"`
	Price     float32 `json:"price"`
	Author    string  `json:"author"`
	Publisher string  `json:"publisher"`
}
