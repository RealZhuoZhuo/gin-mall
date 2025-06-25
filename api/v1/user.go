package v1

import (
	"github.com/RealZhuoZhuo/gin-mall/pkg/e"
	"github.com/RealZhuoZhuo/gin-mall/serializer"
	"github.com/RealZhuoZhuo/gin-mall/service"
	"github.com/gin-gonic/gin"
)

func UserRegisterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userService := service.UserService{}
		if err := c.ShouldBind(&userService); err != nil {
			c.JSON(200, serializer.Response{
				Status: 500,
				Msg:    e.GetMsg(500),
				Error:  err.Error(),
			})
		}
		res := userService.Register()
		c.JSON(200, res)
	}
}
func UserLoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userService := service.UserService{}
		err := c.ShouldBind(&userService)
		if err != nil {
			c.JSON(200, serializer.Response{
				Status: 500,
				Msg:    e.GetMsg(500),
				Error:  err.Error(),
			})
		}
		res := userService.Login()
		c.JSON(200, res)
	}
}
