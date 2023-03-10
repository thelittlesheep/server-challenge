package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	var myEnv map[string]string
	myEnv, _ = godotenv.Read()

	router := gin.Default()

	router.Run(":" + myEnv["PORT"])
}
