package helper

import (
	"github.com/golang-jwt/jwt"
	. "serverChallenge/config"
	. "serverChallenge/models"
	"time"
)

type authClaims struct {
	jwt.StandardClaims
	UserID uint `json:"userId"`
}

func GenerateToken(user User) (string, error) {
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
