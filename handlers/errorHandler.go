package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandler(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			// TODO add logger
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Server Error",
			})
		}

		statusCode := ctx.Writer.Status()
		err := ctx.Errors.Last()

		switch statusCode {
		case 401:
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
		case 403:
			ctx.JSON(http.StatusForbidden, gin.H{
				"message": "Forbidden",
			})
		case 404:
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "Not Found",
			})
		default:
			ctx.JSON(statusCode, gin.H{
				"message": err,
			})
		}
		ctx.Abort()
	}()

	ctx.Next()
}
