package handlers

import (
	"github.com/gin-gonic/gin"
	. "serverChallenge/helper"
	. "serverChallenge/models"
	"strings"
)

func VerifyUser(ctx *gin.Context) {
	ctx.Status(422)

	tokenString := strings.TrimPrefix(ctx.GetHeader("Authorization"), "bearer ")

	claim, err := ValidateToken(tokenString)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Abort()
	}
	if claim.UserID == 0 {
		_ = ctx.Error(err)
		ctx.Abort()
	}

	user := &User{}
	userFromDB, err := user.FindOne(claim.UserID)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Abort()
	}

	ctx.Set("auth", AuthInContext{
		User: *userFromDB,
	})

	ctx.Next()
}
