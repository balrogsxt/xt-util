package xstr

import (
	"math/rand"
	"strings"
)

// Rand 随机长度字符串
func Rand(size int, strs ...string) string {
	str := "abcdfeghjiklnmopqrstuvwxyz0123456789"
	if len(strs) > 0 {
		str = strs[0]
	}
	s := ""
	list := strings.Split(str, "")
	for i := 0; i < size; i++ {
		k := list[rand.Intn(len(list)-1)]
		s += k
	}
	return s
}
