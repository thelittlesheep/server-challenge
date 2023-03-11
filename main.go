package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"serverChallenge/handlers"
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

	router.Use(handlers.ErrorHandler)

	router.Run(":" + myEnv["PORT"])
}
