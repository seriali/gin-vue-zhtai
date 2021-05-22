package message

const (
	SUCCESS           = 200
	ERROR             = 500
	UsernameNotExit   = 1001
	UserPasswordWrong = 1002
	TokenNotExit      = 1003
	TokenRunTime      = 1004
	TokenWrong        = 1005
	TokenTypeWrong    = 1006
	LoginFail         = 1007
)

var Code = map[int]string{
	SUCCESS:           "success",
	ERROR:             "failed",
	UsernameNotExit:   "用户名不存在",
	UserPasswordWrong: "密码错误",
	TokenNotExit:      "token不存在",
	TokenRunTime:      "token已过期",
	TokenWrong:        "token不正确",
	TokenTypeWrong:    "token格式错误",
	LoginFail:         "登录失败",
}

func GetMsg(code int) string {
	return Code[code]
}
