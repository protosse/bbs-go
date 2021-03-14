package common

var TipErrorCode = 1000

type CodeError struct {
	Code    int
	Message string
}

var (
	CaptchaError       = &CodeError{1001, "验证码错误"}
	EmailExistError    = &CodeError{1002, "邮箱已被占用"}
	CreateUserError    = &CodeError{1003, "创建用户失败"}
	ParamError         = &CodeError{1004, "请求参数错误"}
	UsernameExistError = &CodeError{1005, "用户名以存在"}
)
