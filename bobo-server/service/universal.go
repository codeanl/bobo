package service

import "bobo-server/dao"

const (
	KEY_CODE               = "code:"              // 验证码
	KEY_USER               = "user:"              // 记录用户
	KEY_ARTICLE_VIEW_COUNT = "article_view_count" // 文章查看数
	KEY_DAILY_VIEW_COUNT   = "daily_view_count"   // 动态查看数

)

var (
	dailyDao   dao.Daily
	articleDao dao.Article
)
