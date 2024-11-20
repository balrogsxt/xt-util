package xstr

import (
	"github.com/balrogsxt/xt-util/standard/xrand"
)

// Rand 随机长度字符串
// Deprecated: 此函数已弃用,换xrand.String替代
func Rand(size int, strs ...string) string {
	return xrand.String(size, strs...)
}
