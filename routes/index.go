package routes

import (
	"github.com/gin-gonic/gin"
	. "serverChallenge/handlers"
)

func Main(router *gin.Engine) {
	indexRouting := router.Group("/")
	{
		indexRouting.GET("", GetIndexView)
	}
	index := router.Group("/api")
	UserRoute(index)
	CommonRoute(index)
}
