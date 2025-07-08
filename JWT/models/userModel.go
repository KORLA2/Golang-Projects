package models

import "time"

type User struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	Name          string    `json:"name" validate:"required,min=3,max=100" `
	Email         string    `json:"email" validate:"required,email" gorm:"not null;uniqueIndex"`
	Password      string    `json:"password" validate:"required,min=6"`
	Phone         string    `json:"phone" validate:"required,numeric" gorm:"not null;uniqueIndex"`
	CreatedAt     time.Time `json:"createdat"`
	UpdatedAt     time.Time `json:"updatedat"`
	UserID        string    `json:"userID" gorm:"uniqueIndex" `
	Token         string    `json:"token"`
	Refresh_Token string    `json:"refresh_token"`
	User_Type     string    `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
}
