package valid

import (
	"regexp"
	"strings"
)

// IsRegex 正则判断
func IsRegex(pattern string, value string) bool {
	has, _ := regexp.MatchString(pattern, value)
	return has
}

// IsEmpty 是否字符串为空
func IsEmpty(s string) bool {
	return len(strings.Trim(s, " ")) == 0
}

// IsEmail 判断是否是邮箱
func IsEmail(s string) bool {
	has, _ := regexp.MatchString("^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$", s)
	return has
}

// IsChineseText 是否全部由中文汉字组成
func IsChineseText(str string) bool {
	return IsRegex("^[\u4e00-\u9fa5\\s]+$", str)
}
