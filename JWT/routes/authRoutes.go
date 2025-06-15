package routes

import "github.com/gin-gonic/gin"

func AuthRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.POST("/user/signup", controllers.SignUp())
	incomingRoutes.POST("/user/signin", controllers.SignIn())
}
