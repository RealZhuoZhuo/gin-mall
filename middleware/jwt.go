package middleware

import (
	"github.com/RealZhuoZhuo/gin-mall/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(
				http.StatusUnauthorized,
				gin.H{
					"error": "没有jwt",
				})
			ctx.Abort()
			return
		}
		username, err := utils.CheckJwt(token)
		if err != nil {
			ctx.JSON(
				http.StatusUnauthorized,
				gin.H{
					"error": err.Error(),
				})
			ctx.Abort()
			return
		}
		ctx.Set("username", username)
		ctx.Next()
	}
}
