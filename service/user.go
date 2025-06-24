package service

import (
	"github.com/RealZhuoZhuo/gin-mall/dao"
	"github.com/RealZhuoZhuo/gin-mall/model"
	"github.com/RealZhuoZhuo/gin-mall/pkg/e"
	"github.com/RealZhuoZhuo/gin-mall/serializer"
)

type UserService struct {
	UserName string
	Password string
	NickName string
	key      string
}

func (service UserService) Register(username string) serializer.Response {
	code := e.SUCCESS
	userDao := dao.UserDao{}
	if exist, err := userDao.UserExistOrNot(username); !exist && err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	} else {
		if err := userDao.Create(&model.User{}); err != nil {
			code = e.ERROR
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
}
