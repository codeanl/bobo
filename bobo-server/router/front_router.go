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
	r.SetTrustedProxies([]string{"*"})     // 设置受信任代理，信任所有代理
	r.Use(middleware.Logger())             // 自定义的 zap 日志中间件
	r.Use(middleware.ErrorRecovery(false)) // 自定义错误处理中间件
	r.Use(middleware.Cors())               // 跨域中间件
	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	// 设置静态文件目录
	r.Static("/uploads", "./uploads")
	// 无需鉴权的接口
	base := r.Group("/api")
	{
		base.POST("/upload", userApi.Upload) // 上传图片
		base.GET("/code", userApi.SendCode)  // 验证码
		//用户
		user := base.Group("/user")
		{
			user.POST("/login", userApi.Login)       // 用户登陆
			user.POST("/register", userApi.Register) // 用户注册
			user.POST("/retrieve", userApi.Retrieve) // 找回账户
			user.GET("/info", userApi.Info)          // 用户详情
		}
		// 贴子
		Daily := base.Group("/daily")
		{
			Daily.GET("/list", dailyApi.DailyList) // 贴子列表
			Daily.GET("/info", dailyApi.DailyInfo) // 贴子详情
		}
		//文章
		Article := base.Group("/article")
		{
			Article.GET("/list", articleApi.ArticleList) // 文章列表
			Article.GET("/info", articleApi.ArticleInfo) // 文章详情
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
		Daily := auth.Group("/daily")
		{
			Daily.POST("/save_or_update", dailyApi.SaveOrUpdate) // 创建/编辑贴子
			Daily.DELETE("/delete", dailyApi.Delete)             // 删除贴子
		}
		//文章
		Article := auth.Group("/article")
		{
			Article.POST("/save_or_update", articleApi.SaveOrUpdate) // 创建/编辑贴
			Article.DELETE("/delete", articleApi.Delete)             // 删除
		}
	}
	return r
}
