package router

import (
	"bobo-server/config"
	"bobo-server/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FrontRouter() http.Handler {
	gin.SetMode(config.Cfg.Server.AppMode)
	r := gin.New()
	r.Use(middleware.Logger())             // 自定义的 zap 日志中间件
	r.Use(middleware.ErrorRecovery(false)) // 自定义错误处理中间件
	//todo 操作日志
	r.Use(middleware.Cors())           // 跨域中间件
	r.SetTrustedProxies([]string{"*"}) // 设置受信任代理，信任所有代理
	// 无需鉴权的接口
	base := r.Group("/api")
	{
		base.GET("/code", userApi.SendCode) // 验证码
		//用户
		user := base.Group("/user")
		{
			user.POST("/login", userApi.UserLogin) // 用户
		}
	}
	// 需要鉴权的接口
	//auth := r.Group("/api")

	r.Run(config.Cfg.Server.FrontPort)
	return r
}
