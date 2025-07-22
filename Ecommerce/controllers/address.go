package controllers

import (
	"myapp/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddAddress(db *gorm.DB) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		userID := ctx.Query("userID")
		var Address models.Address
		ctx.BindJSON(Address)
		Address.UserID = userID
		var count int
		if err := db.Table("address").
			Select("COALESCE(MAX(ano), 0)").
			Where("address.user_id=?", userID).
			Scan(&count).Error; err != nil {

			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"Unable to find address": err,
			})
			return
		}

		Address.ANo = count + 1

		if err := db.Create(&Address).Error; err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"Unable to Create Address": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"Success": Address,
		})

	}

}

func DeleteAddress(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		UID := ctx.Query("UID")
		AID, err2 := strconv.Atoi(ctx.Query("ANo"))

		if UID == "" || err2 != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"UserID or AddressID Cannot be null": "",
			})
		}

		target := models.Address{
			ANo:    AID,
			UserID: UID,
		}

		if err := db.Unscoped().Where(&target).Delete(&models.Address{}).Error; err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"Failed to delete the Address": err,
			})
		}
		ctx.JSON(200, gin.H{
			"Successfully Deleted the address for ": target,
		})

	}

}

func EditAddress() gin.HandlerFunc {

	return func(ctx *gin.Context) {

	}

}
