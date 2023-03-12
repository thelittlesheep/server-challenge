package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetIndexView(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}
