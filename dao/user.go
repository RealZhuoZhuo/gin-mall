package dao

import (
	"errors"
	"github.com/RealZhuoZhuo/gin-mall/model"
	"github.com/RealZhuoZhuo/gin-mall/pkg/utils"
	"gorm.io/gorm"
	"log"
)

type UserDao struct {
	*gorm.DB
}

func GetUserDao() *UserDao {
	// 确保 _db 被正确初始化
	if _db == nil {
		log.Fatal("数据库连接未初始化")
	}
	return &UserDao{_db}
}
func (dao *UserDao) UserExistOrNot(userName string) bool {
	var cnt int64
	cnt = 0
	dao.DB.Model(model.User{}).Where("user_name = ?", userName).Count(&cnt)
	if cnt > 0 {
		return true
	}
	return false
}
func (dao *UserDao) Create(user *model.User) error {
	return dao.DB.Create(user).Error
}
func (dao *UserDao) VaildPwd(username string, pwd string) (bool, error) {
	var user model.User
	dao.DB.Model(&model.User{}).Where("user_name = ?", username).First(&user)
	hashpwd := user.PasswordDigest
	if ok := utils.CheckPassword(pwd, hashpwd); ok {
		return true, nil
	}
	return false, errors.New("密码错误")
}
