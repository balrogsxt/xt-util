package xtime

import (
	"fmt"
	"time"
)

// MonthFirst 获取本月第一天的时间
func (l *Time) MonthFirst() *Time {
	t := time.Date(l.Year(), l.Month(), 1, 0, 0, 0, 0, l.Location())
	return New(t)
}

// MonthLast 获取本月最后一天时间
func (l *Time) MonthLast() *Time {
	t := time.Date(l.Year(), l.Month()+1, 1, 23, 59, 59, 999, l.Location())
	return New(t.AddDate(0, 0, -1))
}

// TodayFirst 获取今日凌晨开始时间
func (l *Time) TodayFirst() *Time {
	layout := fmt.Sprint(time.DateOnly, "00:00:00")
	t, _ := time.ParseInLocation(layout, time.Now().Format(layout), time.Local)
	return New(t)
}

// TodayLast 获取今天晚上最后时间
func (l *Time) TodayLast() *Time {
	return New(time.Unix(l.TodayFirst().Unix()+86400-1, 0))
}

// WeekFirst 本周第一天时间
func (l *Time) WeekFirst(weeks ...int) *Time {
	week := 0
	if len(weeks) > 0 {
		week = weeks[0]
	}
	now := l.wrapper
	weekday := now.Weekday()
	var daysToSubtract int
	if weekday == time.Sunday { //如果是星期天
		daysToSubtract = 6
	} else {
		daysToSubtract = int(weekday) - 1
	}
	d := now.AddDate(0, 0, -daysToSubtract+week*7)
	return New(time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location()))
}

// WeekLast 获取本周最后一天时间
func (l *Time) WeekLast(weeks ...int) *Time {
	t := l.WeekFirst(weeks...).AddDate(0, 0, 7).Add(-time.Millisecond)
	return New(t)
}
