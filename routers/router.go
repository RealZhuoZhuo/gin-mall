package routers

import (
	api "github.com/RealZhuoZhuo/gin-mall/api/v1"
	"github.com/RealZhuoZhuo/gin-mall/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	auth := r.Group("api/v1")
	{
		auth.GET("/ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		auth.POST("/user/register", api.UserRegisterHandler())
		auth.POST("/user/login", api.UserLoginHandler())
	}
	upload := r.Group("api/v1")
	upload.Use(middleware.AuthMiddleware())
	{
		upload.POST("/user/avatar", api.UploadAvatarHandler())
	}
	return r
}
