package util

import (
	"errors"
	"strings"

	"regexp"
)

func IsUsername(username string) (err error) {
	str := strings.TrimSpace(username)
	if len(str) == 0 {
		err = errors.New("请输入用户名")
		return
	}

	pattern := `^[a-zA-Z][0-9a-zA-Z_-]{4,11}$`
	matched, err := regexp.MatchString(pattern, str)
	if !matched {
		err = errors.New("用户名必须由5-12位(数字、字母、_、-)组成，且必须以字母开头")
		return
	}
	return
}

func IsEmail(email string) (err error) {
	str := strings.TrimSpace(email)
	if len(str) == 0 {
		err = errors.New("请输入邮箱")
		return
	}
	pattern := `^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$`
	matched, _ := regexp.MatchString(pattern, str)
	if !matched {
		err = errors.New("邮箱格式不符合规范")
	}
	return
}

func IsPassword(password string) (err error) {
	str := strings.TrimSpace(password)
	if len(str) == 0 {
		err = errors.New("请输入密码")
		return
	}

	if len(str) < 6 {
		err = errors.New("密码不能小于6位")
		return
	}
	return
}
