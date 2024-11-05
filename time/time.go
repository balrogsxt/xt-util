package ytime

import (
	"bytes"
	"fmt"
	ystr "github.com/balrogsxt/yuki-tools/str"
	yvar "github.com/balrogsxt/yuki-tools/var"
	"strconv"
	"strings"
	"time"
)

type Time struct {
	t time.Time
}

func Now() *Time {
	return New()
}
func New(ts ...any) *Time {
	var t time.Time
	if len(ts) > 0 {
		switch tp := ts[0].(type) {
		case string: //字符串时间转换
			for _, layout := range formatLayouts {
				if tt, err := time.ParseInLocation(layout, tp, time.Local); err == nil {
					return newTime(tt, nil)
				}
			}

		case time.Time:
			return newTime(tp, nil)
		case uint, uint32, uint64, int, int32, int64:
			//判断是否是毫秒级时间戳
			timestamp := yvar.New(tp).Int64()
			if timestamp > 1e12 {
				t = time.Unix(timestamp/1000, timestamp%1000)
			} else {
				t = time.Unix(timestamp, 0)
			}
			return newTime(t, time.Local)
		default:
			return &Time{
				t: time.Time{},
			}
		}
	}
	return &Time{
		t: time.Now(),
	}
}
func newTime(t time.Time, local *time.Location) *Time {
	if local != nil {
		//设置指定时区
		t = t.In(local)
	}
	return &Time{
		t: t,
	}
}

func (l *Time) String() string {
	if l.t.IsZero() {
		return ""
	}
	return l.t.Format("2006-01-02 15:04:05")
}

// Time 获取原生time.Time
func (l *Time) Time() time.Time {
	return l.t
}

// Format 格式化时间,支持部分php常见格式化
func (l *Time) Format(format string) string {
	//支持按照php方式调用Y-m-d H:i:s
	strs := strings.Split(format, "")
	buffer := bytes.NewBuffer(nil)
	for i := 0; i < len(strs); i++ {
		switch strs[i] {
		case "Y": //年份的四位数表示
			buffer.WriteString(fmt.Sprintf("%04d", l.t.Year()))
		case "m": //月份的数字表示（从 01 到 12）
			buffer.WriteString(fmt.Sprintf("%02d", l.t.Month()))
		case "d": //一个月中的第几天（从 01 到 31）
			buffer.WriteString(fmt.Sprintf("%02d", l.t.Day()))
		case "H": //24 小时制，带前导零（00 到 23）
			buffer.WriteString(fmt.Sprintf("%02d", l.t.Hour()))
		case "i": //分，带前导零（00 到 59）
			buffer.WriteString(fmt.Sprintf("%02d", l.t.Minute()))
		case "s": //秒，带前导零（00 到 59）
			buffer.WriteString(fmt.Sprintf("%02d", l.t.Second()))
		case "D": //星期几的文本表示（用三个字母表示）
			buffer.WriteString(ystr.Substr(l.t.Weekday().String(), 0, 3))
		case "l": //星期几的完整的文本表示
			buffer.WriteString(l.t.Weekday().String())
		case "j": //一个月中的第几天，不带前导零（1 到 31）
			buffer.WriteString(strconv.Itoa(l.t.Day()))
		case "w": //星期的数字表示0123456,0代表星期天
			buffer.WriteString(strconv.Itoa(int(l.t.Weekday())))
		case "N": //星期的数字表示1234567,7代表星期天
			n := int(l.t.Weekday())
			if n == 0 {
				n = 7
			}
			buffer.WriteString(strconv.Itoa(n))
		case "z": //一年中的第几天（从 0 到 365）
			buffer.WriteString(strconv.Itoa(l.t.YearDay()))
		case "F": //月份的完整的文本表示（January[一月份] 到 December[十二月份]）
			buffer.WriteString(l.t.Month().String())
		case "M": //月份的短文本表示（用三个字母表示）
			buffer.WriteString(ystr.Substr(l.t.Month().String(), 0, 3))
		case "n": //月份的数字表示，不带前导零（1 到 12）
			buffer.WriteString(strconv.Itoa(int(l.t.Month())))
		case "t": //给定月份中包含的天数
			d := time.Date(l.t.Year(), l.t.Month(), 1, 0, 0, 0, 0, l.t.Location()).AddDate(0, 1, -1).Day()
			buffer.WriteString(strconv.Itoa(d))
		case "o": // ISO-8601 标准下的年份数字
			buffer.WriteString(strconv.Itoa(l.t.Year()))
		case "a": //小写形式表示：am 或 pm
			buffer.WriteString(l.t.Format("pm"))
		case "A": //大写形式表示：AM 或 PM
			buffer.WriteString(l.t.Format("PM"))
		case "g": //12 小时制，不带前导零（1 到 12）
			buffer.WriteString(l.t.Format("3"))
		case "G": //24 小时制，不带前导零（0 到 23）
			buffer.WriteString(strconv.Itoa(l.t.Hour()))
		case "h": //12 小时制，带前导零（01 到 12）
			buffer.WriteString(l.t.Format("03"))
		case "e": // 时区标识符（例如：UTC、GMT、Atlantic/Azores）
			buffer.WriteString(l.t.Location().String())
			//case "L": //是否是闰年（如果是闰年则为 1，否则为 0）
		//case "u": //微秒
		//	buffer.WriteString(strconv.Itoa(int(l.t.UnixNano() / 1000)))
		//case "I": //日期是否是在夏令时（如果是夏令时则为 1，否则为 0）
		//case "O": //格林威治时间（GMT）的差值，单位是小时（实例：+0100）
		//case "P": // 格林威治时间（GMT）的差值，单位是 hours:minutes
		//case "T": //时区的简写（实例：EST、MDT）
		//case "Z": //以秒为单位的时区偏移量。UTC 以西时区的偏移量为负数（-43200 到 50400）
		case "c": //ISO-8601 标准的日期（例如 2013-05-05T16:34:42+00:00）
			buffer.WriteString(l.t.Format(time.RFC3339))
		case "r": //RFC 2822 格式的日期（例如 Fri, 12 Apr 2013 12:01:05 +0200）
			buffer.WriteString(l.t.Format(time.RFC1123Z))
		case "U": //自 Unix 纪元（January 1 1970 00:00:00 GMT）以来经过的秒数
			buffer.WriteString(strconv.Itoa(int(l.t.Unix())))
		default:
			buffer.WriteString(strs[i])
		}
	}
	return buffer.String()
}

// Distance 计算当前时间距离指定的时间过了多久(处理以秒为单位)
func (l *Time) Distance(t time.Time, reverses ...bool) string {
	reverse := false
	if len(reverses) > 0 {
		reverse = reverses[0]
	}
	if reverse {
		return distanceSecond(l.t.Unix() - t.Unix())
	} else {
		return distanceSecond(t.Unix() - l.t.Unix())
	}
}

func distanceSecond(second int64) string {
	var (
		i int64 = 60
		h       = i * 60
		d       = h * 24
	)
	if 0 >= second {
		return ""
	}
	if i > second {
		return fmt.Sprintf("%ds", second)
	} else if second >= i && h > second {
		return fmt.Sprintf("%dm%s", second/i, distanceSecond(second%i))
	} else if second >= h && d > second {
		return fmt.Sprintf("%dh%s", second/h, distanceSecond(second%h))
	} else {
		return fmt.Sprintf("%dd%s", second/d, distanceSecond(second%d))
	}
}
