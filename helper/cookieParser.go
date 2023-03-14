package helper

import "github.com/gin-gonic/gin"

func CookieParser(ctx *gin.Context, keyName string) (cookieValue *string, err error) {
	// 根據cookie名字讀取cookie值
	data, err := ctx.Cookie(keyName)
	if err != nil {
		// 直接返回cookie值
		return nil, err
	}

	return &data, nil
}
