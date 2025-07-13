package routes

import "github.com/gin-gonic/gin"

func NonAuthRoutes(router *gin.Engine) {

	router.POST("/user/signup",controllers.SignUp)
	router.POST("/user/signin",controllers.SignIn)
	router.POST("/admin/addproduct",controllers.AddProduct)
    router.GET("/user/products",controllers.GetProducts);
    router.GET("/user/search",controllers.SearchProducts);


}
func AuthRoutes(router * gin.Engine){

	router.Use(middleware.Authenticate())

	router.GET("/addtocart",controllers.AddToCart())
	router.GET("/removeItem",controllers.RemoveItem())
	router.GET("/instantbuy",controllers.InstantBuy())
	router.GET("/cartcheckout",controllers.CartCheckOut())
}