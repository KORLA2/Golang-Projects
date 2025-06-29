package models

import "time"

type User struct {
	ID       uint
	Name     string    `json:"name" validate:"required" `
	Email    string    `json:"email" validate:"required email" `
	Phone    string    `json:"phone" validate:"required numeric" `
	CreateAt time.Time `json:"createdat"`
	UpdateAt time.Time `json:"updatedat"`
	Userid   string
}
