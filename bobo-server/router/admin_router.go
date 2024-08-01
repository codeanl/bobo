package router

import (
	"bobo-server/config"
	"bobo-server/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminRouter() http.Handler {
	gin.SetMode(config.Cfg.Server.AppMode)
	r := gin.New()
	r.Use(middleware.Cors())           // 跨域中间件
	r.SetTrustedProxies([]string{"*"}) // 设置受信任代理，信任所有代理
	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}
