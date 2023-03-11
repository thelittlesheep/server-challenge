package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"serverChallenge/helper"
	. "serverChallenge/models"
	"serverChallenge/resources"
)

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body User true "User"
// @Success 200 {object} User
// @Router /users/ [post]
func CreateUser(ctx *gin.Context) {
	ctx.Status(422)

	user := &User{}
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	err = validateUserDuplicate(user.Email)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	user.Password = string(hashPass)
	user.Create()

	ctx.JSON(http.StatusOK, helper.WrapResponse(resources.Message{Message: "User created"}))
}

func validateUserDuplicate(email string) error {
	user := &User{}

	_, err := user.FindOneByEmail(email)

	if err == nil {
		return errors.New("User already exists")
	}

	return nil
}
