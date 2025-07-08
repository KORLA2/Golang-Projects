package middleware

import (
	"net/http"

	"github.com/Goutham/Gin/helper"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("token")
		if token == "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "token-not-found-in header",
			})
			return
		}

		claims, err := helper.VerifyToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"User Not Authorized": err,
			})
		}
		ctx.Set("email", claims.Email)
		ctx.Set("Name", claims.Name)
		ctx.Set("uid", claims.UserID)
		ctx.Set("user-type", claims.User_Type)
	}

}
