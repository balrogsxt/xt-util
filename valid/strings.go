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
