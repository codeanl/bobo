package router

import (
	"bobo-server/api/front"
	"bobo-server/config"
	"log"
	"net/http"
	"time"
)

var (
	userApi    front.User
	dailyApi   front.Daily
	articleApi front.Article
)

// 后台服务
func AdminServer() *http.Server {
	backPort := config.Cfg.Server.AdminPort
	log.Printf("后台服务启动于 %s 端口", backPort)
	return &http.Server{
		Addr:         backPort,
		Handler:      AdminRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

// 前台服务
func FrontServer() *http.Server {
	frontPort := config.Cfg.Server.FrontPort
	log.Printf("前台服务启动于 %s 端口", frontPort)
	return &http.Server{
		Addr:         frontPort,
		Handler:      FrontRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
