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
	UsernameExistError = &CodeError{1004, "用户名已存在"}
	UserNotExistError  = &CodeError{1005, "用户不存在"}
	PasswordError      = &CodeError{1006, "密码错误"}
	TokenError         = &CodeError{1007, "token错误"}
)
