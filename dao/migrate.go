package dao

import "github.com/RealZhuoZhuo/gin-mall/model"

func migrate() {
	_db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(
		model.Address{}, model.Admin{}, model.Cart{}, model.Category{}, model.User{}, model.Product{}, model.Order{},
	)
}
