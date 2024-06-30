package config

var Cfg Config

// 配置文件的结构体
type Config struct {
	Server  Server
	Mysql   Mysql
	Redis   Redis
	Zap     Zap
	Captcha Captcha
	Email   Email
}

type Server struct {
	AppMode   string
	BackPort  string
	FrontPort string
}

type Mysql struct {
	Host     string // 服务器地址
	Port     string // 端口
	Config   string // 高级配置
	Dbname   string // 数据库名
	Username string // 数据库用户名
	Password string // 数据库密码
	LogMode  string // 日志级别
}

type Redis struct {
	DB       int    // 指定 Redis 数据库
	Addr     string // 服务器地址:端口
	Password string // 密码
}

type Zap struct {
	Level        string // 级别
	Prefix       string // 日志前缀
	Format       string // 输出
	Directory    string // 日志文件夹
	MaxAge       int    // 日志留存时间
	ShowLine     bool   // 显示行
	LogInConsole bool   // 输出控制台
}

type Captcha struct {
	SendEmail  bool // 是否通过邮箱发送验证码
	ExpireTime int  // 过期时间
}

type Email struct {
	To       string // 收件人 多个以英文逗号分隔 例：a@qq.com,b@qq.com
	From     string // 发件人 要发邮件的邮箱
	Host     string // 服务器地址, 例如 smtp.qq.com 前往要发邮件的邮箱查看其 smtp 协议
	Secret   string // 密钥, 不是邮箱登录密码, 是开启 smtp 服务后获取的一串验证码
	Nickname string // 发件人昵称, 通常为自己的邮箱名
	Port     int    // 前往要发邮件的邮箱查看其 smtp 协议端口, 大多为 465
	IsSSL    bool   // 是否开启 SSL
}
