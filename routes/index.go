package routes

import (
	"github.com/gin-gonic/gin"
)

func Main(router *gin.Engine) {
	index := router.Group("/api")
}
