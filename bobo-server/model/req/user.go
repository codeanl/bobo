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

type UpdateProfileReq struct {
	Nickname string `json:"nickname"  label:"昵称"`
	Avatar   string `json:"avatar" label:"头像"`
}

type UpdatePasswordReq struct {
	OldPassword       string `json:"old_password" validate:"required" label:"旧密码"`
	NewPassword       string `json:"new_password" validate:"required" label:"新密码"`
	SecondNewPassword string `json:"second_new_password" validate:"required" label:"确认新密码"`
}

type UpdateEmailReq struct {
	Email string `json:"email" validate:"required" label:"邮箱"`
	Code  string `json:"code" validate:"required" label:"验证码"`
}

type RetrieveReq struct {
	Email    string `json:"email" validate:"required" label:"邮箱"`
	Code     string `json:"code" validate:"required" label:"验证码"`
	Password string `json:"password" validate:"required" label:"密码"`
}
