package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	. "serverChallenge/helper"
	. "serverChallenge/models"
	. "serverChallenge/resources/user"
	"strings"
)

func GetIndexView(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func GetLoginView(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", nil)
}

func GetDashboardView(ctx *gin.Context) {
	errorToErrorPage := errors.New("You are not logged in")
	token, err := CookieParser(ctx, "token")
	if token == nil {
		ctx.Status(http.StatusUnauthorized)
		_ = ctx.Error(errorToErrorPage)
		return
	}

	tokenString := strings.TrimPrefix(*token, "bearer ")

	claim, err := ValidateToken(tokenString)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	user := &User{}
	userFromDB, err := user.FindOne(claim.UserID)
	usersInDetail := []DetailResource{
		{ID: userFromDB.ID, Name: userFromDB.Name, Email: userFromDB.Email},
	}
	if err != nil {
		ctx.Status(http.StatusNotFound)
		_ = ctx.Error(err)
		return
	}

	ctx.HTML(http.StatusOK, "dashboard.html", gin.H{
		"usersInDetail": usersInDetail,
	})
}
