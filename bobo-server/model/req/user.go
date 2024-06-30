package req

type UserLoginReq struct {
	Username string `json:"username" validate:"required" label:"用户名"`
	Password string `json:"password" validate:"required" label:"密码"`
}

type UserRegisterReq struct {
	Username string `json:"username" validate:"required" label:"用户名"`
	Password string `json:"password" validate:"required" label:"密码"`
	Nickname string `json:"nickname" validate:"required" label:"昵称"`
	Email    string `json:"email" validate:"required" label:"邮箱"`
	Code     string `json:"code" validate:"required" label:"验证码"`
}
