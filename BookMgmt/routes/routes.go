package routes

import (
	"github.com/Goutham/BookMgmt/controller"
	"github.com/Goutham/BookMgmt/storage"
	"github.com/gin-gonic/gin"
)

func SetUproutes(incomingRoutes *gin.Engine, db *storage.Repository) {

	incoming := incomingRoutes.Group("/api")
	incoming.POST("/create-book", controller.CreateBook(db))
	incoming.DELETE("/delete-book/:bookid", controller.DeleteBook(db))
	incoming.GET("/get-book/:bookid", controller.GetBookByID(db))
	incoming.GET("/books", controller.GetAllBooks(db))
	incoming.PUT("/update/:bookid", controller.UpdateBook(db))
}
