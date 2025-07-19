package database

import (
	"myapp/models"

	"gorm.io/gorm"
)

func AddtoCart(DB *gorm.DB, ProductID, UserID int) (*models.CartItem, error) {

	CartItem := &models.CartItem{
		UserID: UserID,
		PID:    ProductID,
	}	

	err := DB.Create(CartItem).Error

	if err != nil {
		return nil, err
	}

	return CartItem, nil
}
