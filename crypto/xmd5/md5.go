package xmd5

import (
	"crypto/md5"
	"fmt"
	"io"
)

func Md5String(v string) string {
	if v != "" {
		h := md5.New()
		_, _ = io.WriteString(h, v)
		return fmt.Sprintf("%x", h.Sum(nil))
	}
	return ""
}

func Md5Byte(v []byte) string {
	if len(v) > 0 {
		h := md5.New()
		h.Write(v)
		return fmt.Sprintf("%x", h.Sum(nil))
	}
	return ""
}
