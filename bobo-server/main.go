package main

import (
	"bobo-server/dao"
	"bobo-server/router"
	"bobo-server/utils"
	"golang.org/x/sync/errgroup"
	"log"
)

var g errgroup.Group

func main() {
	// 初始化 Viper
	utils.InitViper()
	// 初始化 Logger
	utils.InitLogger()
	//// 初始化数据库
	dao.DB = utils.InitMySQL()
	// 初始化 Redis
	utils.InitRedis()
	// 前台接口服务
	g.Go(func() error {
		return router.FrontServer().ListenAndServe()
	})

	// 后台接口服务
	g.Go(func() error {
		return router.AdminServer().ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

}
