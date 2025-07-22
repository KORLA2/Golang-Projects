package models

import (
	"time"
)

type User struct {
	ID            int        `json:"id" gorm:"primaryKey;autoIncrement"`
	Name          string     `json:"name" validate:"required" `
	Phone         string     `json:"phone" validate:"required,numeric,min" gorm:"unique;notnull" `
	Email         string     `json:"email" validate:"required,email" gorm:"unique;notnull"`
	Password      string     `json:"password" validate:"required,min=6"`
	CreatedAt     time.Time  `json:"createdat"`
	UpdateddAt    time.Time  `json:"updatedat"`
	UserID        string     `json:"userID" gorm:"unique;notnull" `
	Token         string     `json:"token"`
	Refresh_Token string     `json:"refresh_token"`
	CartItems     []CartItem `json:"cart" gorm:"foreignKey:UserID"`
	Orders        []Order    `json:"orders" gorm:"foreignKey:UserID"`
	Addresses     []Address  `json:"address" validate:"required" gorm:"foreignKey:UserID"`
}

type Order struct {
	OID        int         `json:"oid" gorm:"primaryKey;autoIncrement"`
	UserID     string      `json:"uid"`
	User       User        `gorm:"foreignKey:UserID;references:UserID"`
	OrderdAt   time.Time   `json:"createdat"`
	Status     string      `json:"status"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	OIID     int     `json:"oiid" gorm:"primaryKey;autoIncrement"`
	OrderID  int     `json:"orderID"`
	Order    Order   `gorm:"foreignKey:OrderID"`
	Quantity int     `json:"quantity" validate:"required,numeric"`
	PID      int     `json:"pid"`
	Product  Product `gorm:"foreignKey:PID"`
}

type CartItem struct {
	CID      int     `json:"cid" gorm:"primaryKey;autoIncrement"`
	UserID   string  `json:"uid"`
	User     User    `gorm:"foreignKey:UserID;references:UserID"`
	Quantity int     `json:"quantity" validate:"required,numeric"`
	PID      int     `json:"pid"`
	Product  Product `gorm:"foreignKey:PID"`
}

type Product struct {
	PID    int     `json:"id" gorm:"primaryKey;autoIncrement"`
	PName  string  `json:"pname" validate:"required"`
	Price  float64 `json:"price" validate:"required,numeric"`
	Image  string  `json:"image" validate:"required"`
	Rating float64 `json:"rating"`
}

type Address struct {
	ID      int    `json:"id" gorm:"primaryKey;autoIncrement"`
	ANo     int    `json:"ano" validate:"required" `
	UserID  string `json:"uid" validate:"required"`
	User    User   `gorm:"foreignKey:UserID;references:UserID"`
	AptName string `json:"aptname" validate:"required" `
	Street  string `json:"street" validate:"required"`
	City    string `json:"city" validate:"required"`
	State   string `json:"state" validate:"required"`
	PinCode string `json:"pincode" validate:"required"`
}
