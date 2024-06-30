package front

import (
	"bobo-server/model/req"
	"bobo-server/utils"
	"bobo-server/utils/r"
	"github.com/gin-gonic/gin"
)

type User struct{}

// SendCode 发送邮件验证码
func (*User) SendCode(c *gin.Context) {
	r.SendCode(c, userService.SendCode(c.Query("email")))
}

// Login 登录
func (*User) Login(c *gin.Context) {
	loginVo, code := userService.Login(c, utils.BindValidJson[req.UserLoginReq](c))
	r.SendData(c, code, loginVo)
}

// Register 注册
func (*User) Register(c *gin.Context) {
	code := userService.Register(c, utils.BindValidJson[req.UserRegisterReq](c))
	r.SendCode(c, code)
}
