package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "serverChallenge/docs"
	"serverChallenge/handlers"
	"serverChallenge/routes"
)

// @title Server Challenge API
// @version 1.0
// @description This is a basic API using Gin and Gorm.
// @host https://lshuang.dev
// @BasePath /api
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

	opts := []func(config *ginSwagger.Config){
		ginSwagger.DefaultModelsExpandDepth(-1)}

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, opts...))

	router.Use(handlers.ErrorHandler)

	router.LoadHTMLGlob("views/*")

	routes.Main(router)

	if myEnv["PORT"] == "" {
		myEnv["PORT"] = "8080"
	}
	router.Run(":" + myEnv["PORT"])
}
