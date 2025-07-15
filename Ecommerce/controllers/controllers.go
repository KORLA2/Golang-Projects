package controller

import (
	"myapp/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var validate = validator.New()

func HashPassword() string {

}

func SignUp(db *gorm.DB) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var user models.User
		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Unable to Bind User": err.Error(),
			})
			return
		}
		if err := validate.Struct(user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Unable to Validate User": err.Error(),
			})
			return
		}

		user.Password = HashPassword()
		user.CreatedAt, _ = time.Parse(time.RFC1123, time.Now().Format(time.RFC1123))
		user.UpdateddAt = user.CreatedAt
		user.Token, user.RefreshToken := helper.GenerateToken(user.Name, user.Phone, user.Email, user.UserID)
		user.CartItems = make([]models.CartItem, 0)
		user.Orders = make([]models.Order, 0)
		user.Addresses = make([]models.Address, 0)

		if err := db.AutoMigrate(&models.User{}); err != nil {

			ctx.JSON(http.StatusInternalServerError, gin.H{
				"user table creation error": err.Error(),
			})
			return
		}

		if err := db.Create(&user).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"user creation failed": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"User SignedUP": user,
		})

	}

}

func SignIn(db *gorm.DB) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var user, foundUser models.User

		if err := ctx.BindJSON(&user); err != nil {

			ctx.JSON(http.StatusBadRequest, gin.H{
				"Unable to Bind User SignIn ": err.Error(),
			})
			return
		}
		if err := db.Where("phone=?", user.Phone).First(&foundUser).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"User Phone Number Doesn't exist": err.Error(),
			})

		}

		token, refreshtoken, _ := helper.GenerateToken(foundUser.Name, foundUser.Phone, foundUser.Email, foundUser.UserID)

		foundUser.Token = token
		foundUser.Refresh_Token = refreshtoken
 
		if err:=verifyPassword(user.Password,foundUser.Password);err!=nil{
			ctx.JSON(http.StatusBadRequest,gin.H{
				"Password Not Correct":err.Error(),
			})
		}


		if err := db.Model(&user).Where("phone=?", user.Phone).Updates(foundUser).Error; err != nil {

			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Unable to Update User": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"Successfully logged in": foundUser,
		})

	}

}
