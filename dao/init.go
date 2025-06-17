package dao

import (
	"github.com/RealZhuoZhuo/gin-mall/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var _db *gorm.DB

func InitDB() {
	dsn := conf.Conf.Mysql.Dns
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.SetConnMaxIdleTime(10)
	sqlDB.SetMaxOpenConns(100)
	_db = db
	migrate()
}
