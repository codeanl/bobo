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
	r.Use(middleware.Cors())               // 跨域中间件
	r.SetTrustedProxies([]string{"*"})     // 设置受信任代理，信任所有代理
	// 无需鉴权的接口
	base := r.Group("/api")
	{
		base.GET("/code", userApi.SendCode) // 验证码
		//用户
		user := base.Group("/user")
		{
			user.POST("/login", userApi.Login)       // 用户登陆
			user.POST("/register", userApi.Register) // 用户注册
		}
		// 贴子
		DailySharing := base.Group("/daily_sharing")
		{
			DailySharing.GET("/list", dailySharingApi.DailySharingList) // 贴子列表
			DailySharing.GET("/info", dailySharingApi.DailySharingInfo) // 贴子详情

		}
	}
	// 需要鉴权的接口
	auth := base.Group("")
	auth.Use(middleware.JWTUserAuth()) // JWT 鉴权中间件
	{
		// 用户模块
		user := auth.Group("/user")
		{
			user.GET("/profile", userApi.Profile)        // 个人详情
			user.PUT("/profile", userApi.UpdateProfile)  // 更新个人信息
			user.PUT("/up_pass", userApi.UpdatePassword) //更换密码
			user.PUT("/up_email", userApi.UpdateEmail)   //更换邮箱
		}
		// 贴子
		DailySharing := auth.Group("/daily_sharing")
		{
			DailySharing.POST("/save_or_update", dailySharingApi.SaveOrUpdate) // 创建/编辑贴子
			DailySharing.DELETE("/delete", dailySharingApi.Delete)             // 删除贴子
		}
	}

	r.Run(config.Cfg.Server.FrontPort)
	return r
}
