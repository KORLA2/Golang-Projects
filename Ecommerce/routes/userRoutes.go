package routes

import (
	controller "myapp/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NonAuthRoutes(router *gin.Engine, DB *gorm.DB) {

	router.POST("/user/signup", controller.SignUp(DB))
	router.POST("/user/signin", controller.SignIn)
	router.POST("/admin/addproduct", controller.AddProduct)
	router.GET("/user/products", controller.GetProducts)
	router.GET("/user/search", controller.SearchProducts)

}
func AuthRoutes(router *gin.Engine) {

	router.Use(middleware.Authenticate())

	router.GET("/addtocart", controller.AddToCart())
	router.GET("/removeItem", controller.RemoveItem())
	router.GET("/instantbuy", controller.InstantBuy())
	router.GET("/cartcheckout", controller.CartCheckOut())
}
