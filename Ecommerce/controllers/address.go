package controllers

import (
	"myapp/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddAddress() {

}

func DeleteAddress(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		UID, err1 := strconv.Atoi(ctx.Query("UID"))
		AID, err2 := strconv.Atoi(ctx.Query("AddressID"))

		if err1 != nil || err2 != nil {
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
