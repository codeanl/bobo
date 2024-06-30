package service

import (
	"bobo-server/config"
	"bobo-server/dao"
	"bobo-server/model"
	"bobo-server/model/dto"
	"bobo-server/model/req"
	"bobo-server/utils"
	"bobo-server/utils/r"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type User struct{}

// SendCode 发送验证码
func (*User) SendCode(email string) (code int) {
	// 已经发送验证码且未过期
	if utils.Redis.GetVal(KEY_CODE+email) != "" {
		return r.ERROR_EMAIL_HAS_SEND
	}
	expireTime := config.Cfg.Captcha.ExpireTime
	validateCode := utils.Encryptor.ValidateCode()
	content := fmt.Sprintf(`
		<div style="text-align:center"> 
			<div style="padding: 8px 40px 8px 50px;">
            	<p>
					您本次的验证码为
					<p style="font-size:75px;font-weight:blod;"> %s </p>
					为了保证账号安全，验证码有效期为 %d 分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用~
				</p>
       	 	</div>
			<div>
            	<p>发送邮件专用邮箱，请勿回复。</p>
        	</div>
		</div>
	`, validateCode, expireTime)
	if err := utils.Email(email, "验证码", content); err != nil {
		return r.ERROR_EMAIL_SEND
	}
	// 将验证码存储到 Redis 中
	utils.Redis.Set(KEY_CODE+email, validateCode, time.Duration(expireTime)*time.Minute)
	return r.OK
}

// Register 注册
func (*User) Register(c *gin.Context, req req.UserRegisterReq) (code int) {
	// 检查用户名是否存在
	if dao.GetOne(model.User{}, "username", req.Username).ID != 0 {
		return r.ERROR_USERNAME_EXIST
	}
	// 检查邮箱是否存在
	if dao.GetOne(model.User{}, "email", req.Email).ID != 0 {
		return r.ERROR_EMAIL_EXIST
	}
	// 检查验证码是否正确
	if utils.Redis.GetVal(KEY_CODE+req.Email) != req.Code {
		return r.ERROR_EMAIL_NOT_EXIST
	}
	// 生成用户信息
	dao.Create(&model.User{
		Username:      req.Username,
		Password:      req.Password,
		Nickname:      req.Username,
		Email:         req.Email,
		Status:        1,
		Role:          2,
		Avatar:        "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT8YhgQ2KuG7qhi2awDluoYJ_tKX7A_aot5Qg&s",
		LastLoginTime: time.Now(),
	})
	return r.OK
}

// Login 登录
func (*User) Login(c *gin.Context, req req.UserLoginReq) (token string, code int) {
	// 检查用户是否存在
	user := dao.GetOne(model.User{}, "username", req.Username)
	if user.ID == 0 {
		return token, r.ERROR_USER_NOT_EXIST
	}
	if user.Status == 0 {
		return token, r.ERROR_USER_DISABLE
	}
	// 检查密码是否正确
	if req.Password != user.Password {
		return token, r.ERROR_PASSWORD_WRONG
	}
	// 获取 IP 相关信息
	ipAddress := utils.IP.GetIpAddress(c)
	ipSource := utils.IP.GetIpSourceSimpleIdle(ipAddress)
	browser, os := "unknown", "unknown"
	if userAgent := utils.IP.GetUserAgent(c); userAgent != nil {
		browser = userAgent.Name + " " + userAgent.Version.String()
		os = userAgent.OS + " " + userAgent.OSVersion.String()
	}
	// 登录信息正确, 生成 Token
	loginInfo := utils.Encryptor.MD5(ipAddress + browser + os) // UUID 生成方法: ip + 浏览器信息 + 操作系统信息
	token, err := utils.GetJWT().GenUserToken(int(user.ID), loginInfo)
	if err != nil {
		utils.Logger.Info("登录时生成 Token 错误: ", zap.Error(err))
		return token, r.ERROR_TOKEN_CREATE
	}
	// 更新用户验证信息: ip 信息 + 上次登录时间
	user.IpAddress = ipAddress
	user.IpSource = ipSource
	user.LastLoginTime = time.Now()
	dao.UpdateOne(user, "id = ?", user.ID)
	// 存入redis
	utils.Redis.Set(KEY_USER+loginInfo, utils.Json.Marshal(dto.UserLoginDTO{
		ID:            int(user.ID),
		CreatedAt:     user.CreatedAt,
		Username:      user.Username,
		Nickname:      user.Nickname,
		Email:         user.Email,
		Avatar:        user.Avatar,
		IpAddress:     user.IpAddress,
		IpSource:      user.IpSource,
		LastLoginTime: user.LastLoginTime,
		Role:          user.Role,
		Token:         token,
	}), time.Duration(config.Cfg.Session.MaxAge)*time.Second)
	return token, r.OK
}
