package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Goutham/Gin/helper"
	"github.com/Goutham/Gin/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var validate = validator.New()

func VerifyPassword(userPassword, foundPassword string) (bool, error) {

	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(foundPassword))

	if err != nil {
		return false, err
	}

	return true, nil
}

func HashPassword(password string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return "", err
	}
	return string(hash), nil

}

func SignIn(db *gorm.DB) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var user models.User
		var foundUser models.User
		if err := ctx.BindJSON(&user); err != nil || user.Phone == "" || user.Password == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Request Not good SignIn": err.Error(),
			})
			return
		}

		if err := db.Where("Phone=?", user.Phone).First(&foundUser).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Phone Number Not found SignIn": err.Error(),
			})
			return
		}

		_, err := VerifyPassword(user.Password, foundUser.Password)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"Password Check Failed": err.Error(),
			})
			return
		}

		token, refresh_token, _ := helper.GenerateAllTokens(foundUser.Email, foundUser.Name, foundUser.UserID, foundUser.User_Type)

		err = helper.UpdateToken(db, token, refresh_token, foundUser.Phone)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"Unable to update the token": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"Log IN successful": foundUser,
		})

	}

}

func SignUp(db *gorm.DB) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var user models.User
		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Request not good Sign Up": err.Error(),
			})
			return
		}

		if validationErr := validate.Struct(user); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Validation Error": validationErr.Error(),
			})
			return
		}
		pass, err := HashPassword(user.Password)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"Cannot Hash Password Sign UP": err.Error(),
			})
		}
		user.Password = pass
		user.UserID = uuid.New().String()
		user.CreatedAt, _ = time.Parse(time.RFC1123, time.Now().Format(time.RFC1123))
		user.UpdatedAt = user.CreatedAt
		token, refresh_token, err := helper.GenerateAllTokens(user.Email, user.Name, user.UserID, user.User_Type)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"Cannot generate Token": err.Error(),
			})
			return
		}

		user.Token = token
		user.Refresh_Token = refresh_token

		if err = db.AutoMigrate(&models.User{}); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Table Creation Failed": err.Error(),
			})
			return
		}

		if err := db.Create(user).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"user": user,
		})

	}

}

func GetUsers(db *gorm.DB) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		if err := helper.CheckUserType(ctx.GetString("user-type")); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		records, _ := strconv.Atoi(ctx.Query("recordperpage"))
		pageno, _ := strconv.Atoi(ctx.Query("pageno"))
		offset := (pageno - 1) * records

		var users = []models.User{}

		if err := db.Offset(offset).Limit(records).Find(&users).Error; err != nil {

			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"Pagination Error": err.Error(),
			})
		}
		ctx.JSON(200, gin.H{
			"users": users,
		})

	}
}

func GetUser(db *gorm.DB) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		userID := ctx.Param("userID")

		if err := helper.MatchUserTypetoUid(ctx, userID); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "You are not allowed to access some body else's page",
			})
			return
		}

		var user models.User

		if err := db.Where("ID=?", userID).First(&user).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Database error",
				"Error":   err.Error(),
			})
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": user,
		})

	}

}
