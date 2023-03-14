package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
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

		matchApiUrl, _ := regexp.MatchString("^/api/.*", ctx.Request.URL.Path)

		if !matchApiUrl && statusCode != 200 {
			ctx.HTML(http.StatusOK, "error.html", gin.H{
				"code":    statusCode,
				"message": ctx.Errors.Last(),
			})
		} else {
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
			}
		}

		ctx.Abort()
	}()

	ctx.Next()
}
