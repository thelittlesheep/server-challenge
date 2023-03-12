package config

import (
	"github.com/joho/godotenv"
	"strconv"
	"time"
)

var myEnv, _ = godotenv.Read()

type AuthConfig struct {
	JWTSecret     string
	JWTExpireTime time.Duration
}

func (AuthConfig) AuthConfig() *AuthConfig {
	// expireTime unit is hour
	expireTime, err := strconv.Atoi(myEnv["JWT_EXPIRE_TIME"])
	if err != nil {
		expireTime = 24
	}
	return &AuthConfig{
		JWTSecret:     myEnv["JWT_SECRET"],
		JWTExpireTime: time.Duration(expireTime),
	}
}
