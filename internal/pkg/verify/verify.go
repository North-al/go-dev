package verify

import (
	"regexp"
)

// 验证是否为邮箱
func IsEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(pattern).MatchString(email)
}

// 验证是否为手机号
func IsPhone(phone string) bool {
	pattern := `^1[3-9]\d{9}$`
	return regexp.MustCompile(pattern).MatchString(phone)
}
