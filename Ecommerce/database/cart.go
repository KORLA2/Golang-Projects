package database

import (
	"errors"
	"myapp/models"

	"gorm.io/gorm"
)

func AddtoCart(DB *gorm.DB, ProductID int, UserID string) (*models.CartItem, error) {

	CartItem := &models.CartItem{
		UserID: UserID,
		PID:    ProductID,
	}

	if err := DB.Where(CartItem).First(&CartItem).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		CartItem.Quantity = 1
		if err := DB.Create(&CartItem).Error; err != nil {
			return nil, err
		}
	} else {
		CartItem.Quantity += 1

		if err := DB.Save(&CartItem).Error; err != nil {
			return nil, err
		}
	}

	return CartItem, nil
}
