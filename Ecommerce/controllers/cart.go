package controllers

import (
	"myapp/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddToCart(DB *gorm.DB) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		ProductParamID := ctx.Param("ProductID")
		UserParamID := ctx.Param("UserID")

		if ProductParamID == "" || UserParamID == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Cannot Add to Cart": "ProductID or UserID is empty",
			})
			return
		}

		ProductID, _ := strconv.Atoi(ProductParamID)
		UserID := UserParamID
		Item, err := database.AddtoCart(DB, ProductID, UserID)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Cannot Add to the Cart": err,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"Successfully Inserted to Cart": Item,
		})

	}

}

func RemoveItemFromCart(DB *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ProductParamID := ctx.Param("ProductID")
		UserParamID := ctx.Param("UserID")

		if ProductParamID == "" || UserParamID == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Cannot Add to Cart": "ProductID or UserID is empty",
			})
			return
		}

		ProductID, _ := strconv.Atoi(ProductParamID)
		UserID, _ := strconv.Atoi(UserParamID)
		Item, err := database.RemoveItemFromCart(DB, ProductID, UserID)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Cannot Remove from the Cart": err,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"Successfully Removed From Cart": Item,
		})

	}
}



func BuyNow(DB *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ProductParamID := ctx.Param("ProductID")
		UserID := ctx.Param("UserID")

		if ProductParamID == "" || UserParamID == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Cannot Buy": "ProductID or UserID is empty",
			})
			return
		}

		ProductID, _ := strconv.Atoi(ProductParamID)
	
		Item, err := database.BuyNow(DB, ProductID, UserID)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Cannot Buy the Product": err,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"Successfully Placed Order": Item,
		})

	}
}

func GetItemFromCart(db *gorm.DB) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		userID := ctx.Query("userID")

		if userID == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"Error": "User ID is required",
			})
			return
		}
		total := 0
		if err := db.Table("cart_items").
			Select("SUM(cart_items.Quantity*products.price)").
			Joins("JOIN products ON cart_items.pid = products.pid").
			Where("cart_items.user_id=?", userID).Scan(&total).Error; err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"Unable to get Items from your Cart": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"Total Value of Items in your Cart": total,
		})

	}

}

