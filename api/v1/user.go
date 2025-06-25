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
		ret := userService.Register()
		c.JSON(200, ret)
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
		ret := userService.Login()
		c.JSON(200, ret)
	}
}
func UploadAvatarHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.GetString("username")
		userService := service.UserService{}
		code := e.SUCCESS
		file, _, err := c.Request.FormFile("avatar")
		if err != nil {
			code = e.ERROR
			c.JSON(200, serializer.Response{
				Status: code,
				Msg:    e.GetMsg(500),
				Error:  err.Error(),
			})
		}
		ret := userService.UploadAvatar(username, file)
		c.JSON(200, ret)
	}
}
