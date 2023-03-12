package helper

import "github.com/gin-gonic/gin"

func WrapResponse(data interface{}) gin.H {
	return gin.H{
		"data": data,
	}
}
