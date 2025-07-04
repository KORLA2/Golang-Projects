package routes

import (
	controller "github.com/Goutham/Gin/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRoutes(incomingRoutes *gin.Engine, db *gorm.DB) {

	incomingRoutes.POST("/user/signup", controller.SignUp(db))
	incomingRoutes.POST("/user/signin", controller.SignIn(db))

}
