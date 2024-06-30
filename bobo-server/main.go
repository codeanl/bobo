package main

import (
	"bobo-server/dao"
	"bobo-server/router"
	"bobo-server/utils"
)

func main() {
	// 初始化 Viper
	utils.InitViper()
	// 初始化 Logger
	utils.InitLogger()
	// 初始化数据库
	dao.DB = utils.InitMySQL()
	// 初始化 Redis
	utils.InitRedis()
	// 初始化前台路由
	router.FrontRouter()
}
