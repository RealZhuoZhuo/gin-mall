package dao

import (
	"github.com/RealZhuoZhuo/gin-mall/model"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func (dao *UserDao) UserExistOrNot(userName string) (bool, error) {
	if err := dao.DB.Model(model.User{}).Where("user_name = ?", userName).Error; err != nil {
		return false, err
	}
	return true, nil
}
func (dao *UserDao) Create(user *model.User) error {
	return dao.DB.Create(user).Error
}
