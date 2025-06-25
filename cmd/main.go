package main

import (
	"github.com/RealZhuoZhuo/gin-mall/conf"
	"github.com/RealZhuoZhuo/gin-mall/dao"
	"github.com/RealZhuoZhuo/gin-mall/routers"
)

func main() {
	conf.InitConfig()
	dao.InitDB()
	r := routers.NewRouter()
	r.Static("static", "static")
	_ = r.Run(":8080")
}
