package front

import (
	"bobo-server/utils/r"
	"github.com/gin-gonic/gin"
)

type User struct{}

// SendCode 发送邮件验证码
func (*User) SendCode(c *gin.Context) {
	r.SendCode(c, userService.SendCode(c.Query("email")))
}
