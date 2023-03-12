package routes

import (
	"github.com/gin-gonic/gin"
	. "serverChallenge/controllers"
)

func CommonRoute(r *gin.RouterGroup) {
	common := r.Group("/common")
	{
		common.POST("/login", func(c *gin.Context) {
			Login(c)
		})
	}
}
