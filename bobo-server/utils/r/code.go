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
	ERROR_EMAIL_SEND     = 002
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
}

func GetMsg(code int) string {
	return codeMsg[code]
}
