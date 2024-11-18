package xtime

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {

	t.Logf(New().WeekLast().Format("Y-m-d H:i:s"))
	t.Logf(New().WeekFirst().Distance(time.Now(), map[string]string{
		"s": "秒",
		"m": "分",
		"h": "时",
		"d": "", //隐藏天
	}))
	t.Logf("%s", time.Now())
	t.Logf("%s", Now())

	t.Log(New(Now().Format("Y-m-d 22:00:00")).Unix())
}

func TestLocalTime(t *testing.T) {
	//mongodb utc time转换为本地时间

	utc := time.Now().In(time.UTC)
	t.Log(utc.Format(time.DateTime))

	t.Log(NewLocal(utc).Format(time.DateTime))

}

func TestLastWeekTime(t *testing.T) {
	t.Logf("上上周开始: %s", Now().WeekFirst(-2).Format(time.DateTime))
	t.Logf("上周开始: %s", Now().WeekFirst(-1).Format(time.DateTime))
	t.Logf("本周开始: %s", Now().WeekFirst().Format(time.DateTime))
	t.Logf("本周结束: %s", Now().WeekLast().Format(time.DateTime))
	t.Logf("下周结束: %s", Now().WeekLast(1).Format(time.RFC3339Nano))
	t.Logf("下下周结束: %s", Now().WeekLast(2).Format(time.RFC3339Nano))
}
