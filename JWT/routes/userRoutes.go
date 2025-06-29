package routes

import (
	controller "github.com/Goutham/Gin/controllers"
	"github.com/Goutham/Gin/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(incomingRoutes *gin.Engine, db *gorm.DB) {

	incomingRoutes.Use(middleware.Auth())
	incomingRoutes.GET("/user/:id", controller.GetUser(db))
	incomingRoutes.GET("/users", controller.GetUsers(db))

}
