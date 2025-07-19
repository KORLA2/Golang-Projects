package controllers

import (
	"myapp/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SearchProducts(db *gorm.DB) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var productList []models.Product

		recordsPerPage, _ := strconv.Atoi(ctx.Query("?recordperpage"))
		pageNo, _ := strconv.Atoi(ctx.Query("?pageNo"))
		offeset := (pageNo - 1) * recordsPerPage

		if err := db.Offset(offeset).Limit(recordsPerPage).Find(&productList).Error; err != nil {

			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Unable to fetch the products": err,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"Success": productList,
		})

	}

}

func SearchProductsByQuery(db *gorm.DB) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		PName:= ctx.Param("ProductName")
		var productList []models.Product
		recordsPerPage, _ := strconv.Atoi(ctx.Query("recordperpage"))
		pageno, _ := strconv.Atoi(ctx.Query("pageno"))
		offset := (pageno - 1) * recordsPerPage

		if err := db.Where("PName=?", PName).Offset(offset).Limit(recordsPerPage).Find(&productList).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Fetching Product by name failed": err,
			})
			return
		}

		if len(productList) == 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"No products Found for the Name": PName,
			})
			return
		}

		ctx.JSON(200, gin.H{
			"Success": productList,
		})
	}

}
