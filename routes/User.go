package routes

import (
	"github.com/gin-gonic/gin"
	. "serverChallenge/controllers"
)

func UserRoute(r *gin.RouterGroup) {
	users := r.Group("/users")
	{
		users.POST("/", func(c *gin.Context) {
			CreateUser(c)
		})
	}
}
