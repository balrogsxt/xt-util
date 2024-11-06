package xtime

import "fmt"

// DistanceSecond 计算一个秒数时间距离过了多久
func DistanceSecond(second int64, formats ...map[string]string) string {
	format := map[string]string{ //可根据参数进行转换语言
		"s": "s", //秒
		"m": "m", //分
		"h": "h", //时
		"d": "d", //天
	}
	if len(formats) > 0 && formats[0] != nil {
		for k, v := range formats[0] {
			format[k] = v
		}
	}

	var (
		i int64 = 60
		h       = i * 60
		d       = h * 24
	)
	if 0 >= second {
		return ""
	}
	if i > second {
		if len(format["s"]) > 0 {
			return fmt.Sprintf("%d%s", second, format["s"])
		} else {
			return ""
		}
	} else if second >= i && h > second {
		if len(format["m"]) > 0 {
			return fmt.Sprintf("%d%s%s", second/i, format["m"], DistanceSecond(second%i, formats...))
		} else {
			return fmt.Sprintf("%s", DistanceSecond(second%i, formats...))
		}
	} else if second >= h && d > second {
		if len(format["h"]) > 0 {
			return fmt.Sprintf("%d%s%s", second/h, format["h"], DistanceSecond(second%h, formats...))
		} else {
			return fmt.Sprintf("%s", DistanceSecond(second%h, formats...))
		}
	} else {
		if len(format["d"]) > 0 {
			return fmt.Sprintf("%d%s%s", second/d, format["d"], DistanceSecond(second%d, formats...))
		} else {
			return fmt.Sprintf("%s", DistanceSecond(second%d, formats...))
		}
	}
}
