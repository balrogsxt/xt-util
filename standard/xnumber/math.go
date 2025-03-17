package xnumber

import (
	"golang.org/x/exp/constraints"
	"math"
)

// RoundFloat 四舍五入到小数点 后 n 位
func RoundFloat(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}

// Max 获取最大的数值
func Max[T constraints.Ordered](slice []T) T {
	var zero T
	if len(slice) == 0 {
		return zero
	}
	var v = slice[0]
	for _, item := range slice {
		if item > v {
			v = item
		}
	}
	return v
}

// Min 取最小值
func Min[T constraints.Ordered](slice []T) T {
	var zero T
	if len(slice) == 0 {
		return zero
	}
	var v = slice[0]
	for _, item := range slice {
		if item < v {
			v = item
		}
	}
	return v
}
