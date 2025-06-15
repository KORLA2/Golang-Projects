package routes

import "github.com/gin-gonic/gin"

func UserRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:id", controllers.GetUser())

}
