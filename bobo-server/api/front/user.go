package front

import (
	"bobo-server/model/req"
	"bobo-server/utils"
	"bobo-server/utils/r"
	"github.com/gin-gonic/gin"
	"strconv"
)

type User struct{}

// Upload 上传文件
func (*User) Upload(c *gin.Context) {
	data, code := utils.Upload(c)
	r.SendData(c, code, data)
}

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
	code := userService.Register(utils.BindValidJson[req.UserRegisterReq](c))
	r.SendCode(c, code)
}

// Profile 个人详情
func (*User) Profile(c *gin.Context) {
	profile, code := userService.Profile(utils.GetFromContext[int](c, "user_id"))
	r.SendData(c, code, profile)
}

// Info 用户详情
func (*User) Info(c *gin.Context) {
	userIdStr := c.Query("user_id")
	userId, _ := strconv.Atoi(userIdStr)
	profile, code := userService.Profile(userId)
	r.SendData(c, code, profile)
}

// UpdateProfile 更新个人信息
func (*User) UpdateProfile(c *gin.Context) {
	r.SendCode(c, userService.UpdateProfile(utils.BindValidJson[req.UpdateProfileReq](c), utils.GetFromContext[int](c, "user_id")))
}

// UpdatePassword 更新密码
func (*User) UpdatePassword(c *gin.Context) {
	r.SendCode(c, userService.UpdatePassword(utils.BindValidJson[req.UpdatePasswordReq](c), utils.GetFromContext[int](c, "user_id")))
}

// UpdateEmail 更新邮箱
func (*User) UpdateEmail(c *gin.Context) {
	r.SendCode(c, userService.UpdateEmail(utils.BindValidJson[req.UpdateEmailReq](c), utils.GetFromContext[int](c, "user_id")))
}

// Retrieve 找回账户
func (*User) Retrieve(c *gin.Context) {
	r.SendCode(c, userService.Retrieve(utils.BindValidJson[req.RetrieveReq](c)))
}
