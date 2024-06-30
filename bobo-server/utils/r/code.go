package r

// 错误码汇总
const (
	OK   = 200
	FAIL = 500

	//通用错误
	ERROR_REQUEST_PARAM = 9001
	ERROR_REQUEST_PAGE  = 9002
	ERROR_INVALID_PARAM = 9003
	ERROR_DB_OPE        = 9004

	//邮箱
	ERROR_EMAIL_HAS_SEND = 8001
	ERROR_EMAIL_SEND     = 8002

	// 用户模块
	ERROR_USER_NAME_USED      = 1001
	ERROR_PASSWORD_WRONG      = 1002
	ERROR_USER_NOT_EXIST      = 1003
	ERROR_USER_NOT_ROLE       = 1004
	ERROR_USER_DISABLE        = 1005
	ERROR_SECOND_PASSWORD_NOT = 1006
	ERROR_USER_EXIST          = 1007
	ERROR_USER_NO_RIGHT       = 1009
	ERROR_OLD_PASSWORD        = 1010
	ERROR_VERIFICATION_CODE   = 1013
	ERROR_EMAIL_NOT_EXIST     = 1014
	ERROR_USERNAME_EXIST      = 1015
	ERROR_EMAIL_EXIST         = 1016
)

var codeMsg = map[int]string{
	OK:   "OK",
	FAIL: "FAIL",

	//通用错误
	ERROR_REQUEST_PARAM: "请求参数格式错误",
	ERROR_REQUEST_PAGE:  "分页参数错误",
	ERROR_INVALID_PARAM: "不合法的请求参数",
	ERROR_DB_OPE:        "数据库操作异常",

	//邮箱
	ERROR_EMAIL_HAS_SEND: "邮箱已发送",
	ERROR_EMAIL_SEND:     "邮件发送失败",

	// 用户模块
	ERROR_USER_NAME_USED:      "用户名已存在",
	ERROR_USER_NOT_EXIST:      "该用户不存在",
	ERROR_PASSWORD_WRONG:      "密码错误",
	ERROR_USER_NO_RIGHT:       "该用户无权限",
	ERROR_OLD_PASSWORD:        "旧密码不正确",
	ERROR_VERIFICATION_CODE:   "验证码错误",
	ERROR_USER_NOT_ROLE:       "用户无角色",
	ERROR_USER_DISABLE:        "账号禁用",
	ERROR_SECOND_PASSWORD_NOT: "两次密码不一致",
	ERROR_USER_EXIST:          "用户已存在",
	ERROR_EMAIL_NOT_EXIST:     "验证码不存在",
	ERROR_USERNAME_EXIST:      "用户名已存在",
	ERROR_EMAIL_EXIST:         "邮箱已存在",
}

func GetMsg(code int) string {
	return codeMsg[code]
}
