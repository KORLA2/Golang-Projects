package routes

import (
	controller "myapp/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NonAuthRoutes(router *gin.Engine, DB *gorm.DB) {

	router.POST("/user/signup", controller.SignUp(DB))
	router.POST("/user/signin", controller.SignIn(DB))
	router.POST("/admin/addproduct", controller.AddProduct(DB))
	router.GET("/user/search", controller.SearchProducts(DB))
	router.GET("/user/search/:ProductID", controller.SearchProductsQuery(DB))

}
func AuthRoutes(router *gin.Engine, DB *gorm.DB) {

	router.Use(middleware.Authenticate())

	router.GET("/addtocart/:ProductID/:UserID", controller.AddToCart(DB))
	router.GET("/removeItem", controller.RemoveItemFromCart(DB))
	router.GET("/buynow", controller.BuyNow(DB))
	router.GET("/cartcheckout", controller.CartCheckOut())
}
