package xtime

import (
	"bytes"
	"fmt"
	"github.com/balrogsxt/xt-util/standard/xvar"
	"github.com/balrogsxt/xt-util/text/xstr"
	"strconv"
	"strings"
	"time"
)

type Time struct {
	wrapper
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
			timestamp := xvar.New(tp).Int64()
			if timestamp > 1e12 {
				t = time.Unix(timestamp/1000, timestamp%1000)
			} else {
				t = time.Unix(timestamp, 0)
			}
			return newTime(t, time.Local)
		default:
			return &Time{
				wrapper{time.Time{}},
			}
		}
	}
	return &Time{
		wrapper{time.Now()},
	}
}
func newTime(t time.Time, local *time.Location) *Time {
	if local != nil {
		//设置指定时区
		t = t.In(local)
	}
	return &Time{
		wrapper{t},
	}
}

func (l *Time) String() string {
	if l.IsZero() {
		return ""
	}
	return l.wrapper.Format(time.DateTime)
}

// Time 获取原生time.Time
func (l *Time) Time() time.Time {
	return l.wrapper.Time
}

// Format 格式化时间,支持部分php常见格式化
func (l *Time) Format(format string) string {
	//支持按照php方式调用Y-m-d H:i:s
	strs := strings.Split(format, "")
	buffer := bytes.NewBuffer(nil)
	for i := 0; i < len(strs); i++ {
		switch strs[i] {
		case "Y": //年份的四位数表示
			buffer.WriteString(fmt.Sprintf("%04d", l.Year()))
		case "m": //月份的数字表示（从 01 到 12）
			buffer.WriteString(fmt.Sprintf("%02d", l.Month()))
		case "d": //一个月中的第几天（从 01 到 31）
			buffer.WriteString(fmt.Sprintf("%02d", l.Day()))
		case "H": //24 小时制，带前导零（00 到 23）
			buffer.WriteString(fmt.Sprintf("%02d", l.Hour()))
		case "i": //分，带前导零（00 到 59）
			buffer.WriteString(fmt.Sprintf("%02d", l.Minute()))
		case "s": //秒，带前导零（00 到 59）
			buffer.WriteString(fmt.Sprintf("%02d", l.Second()))
		case "D": //星期几的文本表示（用三个字母表示）
			buffer.WriteString(xstr.Substr(l.Weekday().String(), 0, 3))
		case "l": //星期几的完整的文本表示
			buffer.WriteString(l.Weekday().String())
		case "j": //一个月中的第几天，不带前导零（1 到 31）
			buffer.WriteString(strconv.Itoa(l.Day()))
		case "w": //星期的数字表示0123456,0代表星期天
			buffer.WriteString(strconv.Itoa(int(l.Weekday())))
		case "N": //星期的数字表示1234567,7代表星期天
			n := int(l.Weekday())
			if n == 0 {
				n = 7
			}
			buffer.WriteString(strconv.Itoa(n))
		case "z": //一年中的第几天（从 0 到 365）
			buffer.WriteString(strconv.Itoa(l.YearDay()))
		case "F": //月份的完整的文本表示（January[一月份] 到 December[十二月份]）
			buffer.WriteString(l.Month().String())
		case "M": //月份的短文本表示（用三个字母表示）
			buffer.WriteString(xstr.Substr(l.Month().String(), 0, 3))
		case "n": //月份的数字表示，不带前导零（1 到 12）
			buffer.WriteString(strconv.Itoa(int(l.Month())))
		case "t": //给定月份中包含的天数
			d := time.Date(l.Year(), l.Month(), 1, 0, 0, 0, 0, l.Location()).AddDate(0, 1, -1).Day()
			buffer.WriteString(strconv.Itoa(d))
		case "o": // ISO-8601 标准下的年份数字
			buffer.WriteString(strconv.Itoa(l.Year()))
		case "a": //小写形式表示：am 或 pm
			buffer.WriteString(l.wrapper.Format("pm"))
		case "A": //大写形式表示：AM 或 PM
			buffer.WriteString(l.Format("PM"))
		case "g": //12 小时制，不带前导零（1 到 12）
			buffer.WriteString(l.Format("3"))
		case "G": //24 小时制，不带前导零（0 到 23）
			buffer.WriteString(strconv.Itoa(l.Hour()))
		case "h": //12 小时制，带前导零（01 到 12）
			buffer.WriteString(l.Format("03"))
		case "e": // 时区标识符（例如：UTC、GMT、Atlantic/Azores）
			buffer.WriteString(l.Location().String())
			//case "L": //是否是闰年（如果是闰年则为 1，否则为 0）
		//case "u": //微秒
		//	buffer.WriteString(strconv.Itoa(int(l.t.UnixNano() / 1000)))
		//case "I": //日期是否是在夏令时（如果是夏令时则为 1，否则为 0）
		//case "O": //格林威治时间（GMT）的差值，单位是小时（实例：+0100）
		//case "P": // 格林威治时间（GMT）的差值，单位是 hours:minutes
		//case "T": //时区的简写（实例：EST、MDT）
		//case "Z": //以秒为单位的时区偏移量。UTC 以西时区的偏移量为负数（-43200 到 50400）
		case "c": //ISO-8601 标准的日期（例如 2013-05-05T16:34:42+00:00）
			buffer.WriteString(l.Format(time.RFC3339))
		case "r": //RFC 2822 格式的日期（例如 Fri, 12 Apr 2013 12:01:05 +0200）
			buffer.WriteString(l.Format(time.RFC1123Z))
		case "U": //自 Unix 纪元（January 1 1970 00:00:00 GMT）以来经过的秒数
			buffer.WriteString(strconv.Itoa(int(l.Unix())))
		default:
			buffer.WriteString(strs[i])
		}
	}
	return buffer.String()
}

// Distance 计算当前时间距离指定的时间过了多久(处理以秒为单位)
func (l *Time) Distance(t time.Time, formats ...map[string]string) string {
	return l.DistanceReverse(t, false, formats...)
}

// DistanceReverse 计算当前时间距离指定的时间过了多久(处理以秒为单位)
func (l *Time) DistanceReverse(t time.Time, reverse bool, formats ...map[string]string) string {
	if reverse {
		return DistanceSecond(l.Unix()-t.Unix(), formats...)
	} else {
		return DistanceSecond(t.Unix()-l.Unix(), formats...)
	}
}
