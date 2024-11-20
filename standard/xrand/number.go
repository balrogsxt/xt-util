package xrand

import (
	"github.com/balrogsxt/xt-util/standard/xnumber"
	"math/rand"
	"time"
)

// Int 返回一个随机数int,随机数包含在min和max在内
func Int(min, max int) int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(max-min+1) + min
}

// Float 返回一个随机数float64,随机数包含在min和max在内
// point 保留小数点位数
func Float(min, max float64, point int) float64 {
	return xnumber.RoundFloat(rand.New(rand.NewSource(time.Now().UnixNano())).Float64()*(max-min)+min, point)
}
