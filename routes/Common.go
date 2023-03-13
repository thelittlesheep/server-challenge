package routes

import (
	"github.com/gin-gonic/gin"
	. "serverChallenge/controllers"
	. "serverChallenge/handlers"
)

func CommonRoute(r *gin.RouterGroup) {
	common := r.Group("/common")
	{
		common.POST("/login", func(c *gin.Context) {
			Login(c)
		})
		common.Use(VerifyUser)
		common.GET("/me", func(c *gin.Context) {
			Me(c)
		})
	}
}
