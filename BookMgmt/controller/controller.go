package controller

import (
	"net/http"

	"github.com/Goutham/BookMgmt/model"
	"github.com/Goutham/BookMgmt/storage"
	"github.com/gin-gonic/gin"
)

func CreateBook(db *storage.Repository) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var book model.Book

		err := ctx.BindJSON(&book)

		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{

				"message": "Could not able to parse body",
			})
			return
		}

		err = db.DB.Create(&book).Error

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Error Storing the Request to the Database",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"book": book,
		})

	}

}

func GetAllBooks(db *storage.Repository) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		books := []model.Book{}
		err := db.DB.Find(&books).Error
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{

				"message": "Could get all books",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": books,
		})

	}

}
func GetBookByID(db *storage.Repository) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		bookid := ctx.Param("bookid")
		Book := model.Book{}

		err := db.DB.Where("id=?", bookid).First(&Book).Error

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Could not able to find the bookid",
			})
			return
		}

		ctx.JSON(200, gin.H{
			"message": "Successfully fecthed the book",
			"book":    Book,
		})

	}

}
func DeleteBook(db *storage.Repository) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		bookid := ctx.Param("bookid")
		var book = model.Book{}
		if bookid == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Book ID cannot be empty",
			})
			return
		}

		err := db.DB.Delete(&book, bookid).Error

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "COuld not able to delete the given bookid",
			})
			return
		}

		ctx.JSON(200, gin.H{
			"message": "successfully deleted the given book ",
		})

	}

}

func UpdateBook(db *storage.Repository) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		Book := model.Book{}
		Result := model.Book{}

		bookid := ctx.Param("bookid")
		ctx.BindJSON(&Result)

		if err := db.DB.First(&Book, bookid).Error; err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Book with given book id not found",
			})
		}

		if err := db.DB.Model(&Book).Updates(Result).Error; err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Cannot update the book",
			})
		}
		ctx.JSON(200, gin.H{
			"data": Result,
		})

	}

}
