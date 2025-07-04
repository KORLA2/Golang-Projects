package helper

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CheckUserType(userType string) (err error) {

	if userType != "ADMIN" {
		return errors.New("you are not admin to call this page")
	}
	return nil

}

func MatchUserTypetoUid(ctx *gin.Context, userID string) (err error) {

	userType := ctx.GetString("user-type")
	uid := ctx.GetString("uid") //need to check once

	if uid != userID && userType == "USER" {
		return errors.New("not allowed to access page")
	}

	return nil
}
