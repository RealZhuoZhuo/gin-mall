package main

import (
	"github.com/RealZhuoZhuo/gin-mall/conf"
	"github.com/RealZhuoZhuo/gin-mall/dao"
)

func main() {
	conf.InitConfig()
	dao.InitDB()
}
