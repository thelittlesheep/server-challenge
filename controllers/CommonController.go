package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"serverChallenge/helper"
	. "serverChallenge/helper"
	. "serverChallenge/models"
	. "serverChallenge/requests/common"
	"serverChallenge/resources/user"
)

// Login godoc
// @Summary Login a user
// @Description Login a user
// @Tags common
// @Accept  json
// @Produce  json
// @Param user body User true "User"
// @Success 200 {object} User
// @Router /login [post]
func Login(ctx *gin.Context) {
	ctx.Status(422)

	user := &LoginSchema{}
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	userFromDB, err := validateUser(user.Email)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), []byte(user.Password))
	if err != nil {
		_ = ctx.Error(errors.New("email or password is incorrect"))
		return
	}

	token, err := helper.GenerateToken(*userFromDB)

	ctx.JSON(http.StatusOK, helper.WrapResponse(token))
}

func validateUser(email string) (*User, error) {
	user := &User{}

	user, err := user.FindOneByEmail(email)

	if err != nil {
		return nil, errors.New("User does not exists")
	}

	return user, nil
}

// Me godoc
// @Summary Get user detail
// @Description Get user detail
// @Tags common
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} User
// @Router /me [get]
func Me(ctx *gin.Context) {
	ctx.Status(422)

	value, exist := ctx.Get("auth")
	if !exist {
		_ = ctx.Error(errors.New("missing auth"))
		return
	}
	auth := value.(AuthInContext)

	userForResponse := user.DetailResource{
		ID:    auth.User.ID,
		Name:  auth.User.Name,
		Email: auth.User.Email,
	}

	ctx.JSON(http.StatusOK, helper.WrapResponse(userForResponse))
}
