package xrand

import (
	"math/rand"
	"strings"
	"time"
)

// String 随机长度字符串
func String(size int, strs ...string) string {
	str := "abcdfeghjiklnmopqrstuvwxyz0123456789"
	if len(strs) > 0 {
		str = strs[0]
	}
	s := ""
	list := strings.Split(str, "")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; i++ {
		k := list[r.Intn(len(list)-1)]
		s += k
	}
	return s
}
