package r

// 错误码汇总
const (
	OK   = 200
	FAIL = 500

	//邮箱
	ERROR_EMAIL_HAS_SEND = 9001
	ERROR_EMAIL_SEND     = 9002
)

var codeMsg = map[int]string{
	OK:   "OK",
	FAIL: "FAIL",

	//邮箱
	ERROR_EMAIL_HAS_SEND: "邮箱已发送",
	ERROR_EMAIL_SEND:     "邮件发送失败",
}

func GetMsg(code int) string {
	return codeMsg[code]
}
