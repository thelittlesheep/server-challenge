package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	. "serverChallenge/config"
	"serverChallenge/helper"
	. "serverChallenge/models"
	. "serverChallenge/requests/common"
	"time"
)

type authClaims struct {
	jwt.StandardClaims
	UserID uint `json:"userId"`
}

func generateToken(user User) (string, error) {
	authConfig := AuthConfig{}.AuthConfig()
	jwtKey := []byte(authConfig.JWTSecret)
	expiresAt := time.Now().Add(authConfig.JWTExpireTime * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, authClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   user.Email,
			ExpiresAt: expiresAt,
		},
		UserID: user.ID,
	})
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

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

	token, err := generateToken(*userFromDB)

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
