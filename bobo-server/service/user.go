package service

import (
	"bobo-server/config"
	"bobo-server/utils"
	"bobo-server/utils/r"
	"fmt"
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
