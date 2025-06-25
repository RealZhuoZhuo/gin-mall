package routers

import (
	api "github.com/RealZhuoZhuo/gin-mall/api/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		v1.POST("/user/register", api.UserRegisterHandler())
		v1.POST("/user/login", api.UserLoginHandler())
	}
	return r
}
