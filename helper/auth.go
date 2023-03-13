package helper

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	. "serverChallenge/config"
	. "serverChallenge/models"
	"time"
)

var authConfig = AuthConfig{}.AuthConfig()
var jwtKey = []byte(authConfig.JWTSecret)

type AuthClaims struct {
	jwt.StandardClaims
	UserID uint `json:"userId"`
}

func GenerateToken(user User) (string, error) {
	expiresAt := time.Now().Add(authConfig.JWTExpireTime * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, AuthClaims{
		StandardClaims: jwt.StandardClaims{
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

func ValidateToken(tokenString string) (*AuthClaims, error) {
	var claims AuthClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return &claims, nil
}
