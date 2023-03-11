package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	var myEnv map[string]string
	myEnv, _ = godotenv.Read()

	var ginMode string
	if myEnv["ENV"] != "production" {
		ginMode = gin.DebugMode
	} else {
		ginMode = gin.ReleaseMode
	}

	gin.SetMode(ginMode)

	router := gin.Default()

	router.Run(":" + myEnv["PORT"])
}
