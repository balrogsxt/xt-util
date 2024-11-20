package xnumber

import (
	"math"
)

// RoundFloat 四舍五入到小数点 后 n 位
func RoundFloat(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}
