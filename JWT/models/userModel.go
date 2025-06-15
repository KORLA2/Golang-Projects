package models

import "time"

type User struct {
	FirstName string `json:"fname" validate: `
	LastName  string
	CreateAt  time.Time
	UpdateAt  time.Time
	userid    string
}
