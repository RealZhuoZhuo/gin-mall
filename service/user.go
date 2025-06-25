package service

import (
	"github.com/RealZhuoZhuo/gin-mall/dao"
	"github.com/RealZhuoZhuo/gin-mall/model"
	"github.com/RealZhuoZhuo/gin-mall/pkg/e"
	"github.com/RealZhuoZhuo/gin-mall/pkg/utils"
	"github.com/RealZhuoZhuo/gin-mall/serializer"
)

type UserService struct {
	UserName string
	Password string
	NickName string
}

func (service UserService) Register() any {
	code := e.SUCCESS
	userDao := dao.GetUserDao()
	if exist := userDao.UserExistOrNot(service.UserName); exist {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "用户存在",
		}
	}
	hashPwd, err := utils.HashPassword(service.Password)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	user := model.User{
		UserName:       service.UserName,
		NickName:       service.NickName,
		PasswordDigest: hashPwd,
	}
	if err := userDao.Create(&user); err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	token, err := utils.GenerateJWT(service.UserName)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.TokenData{
		User:  user,
		Token: token,
	}
}
func (service UserService) Login() any {
	code := e.SUCCESS
	userDao := dao.GetUserDao()
	exist := userDao.UserExistOrNot(service.UserName)
	if !exist {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "用户不存在",
		}
	}
	ok, err := userDao.VaildPwd(service.UserName, service.Password)
	if !ok && err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	token, err := utils.GenerateJWT(service.UserName)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.TokenData{
		User:  nil,
		Token: token,
	}
}
