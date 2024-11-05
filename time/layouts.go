package ytime

import "time"

var formatLayouts = []string{
	time.DateTime,
	time.DateOnly,
	time.TimeOnly,
	time.RFC3339,
	time.RFC3339Nano,
	time.RFC822,
	time.RFC822Z,
	time.RFC850,
	time.RFC1123,
	time.RFC1123Z,
	time.Layout,
	time.ANSIC,
	time.UnixDate,
	time.RubyDate,
	time.Kitchen,
	time.Stamp,
	time.StampMilli,
	time.StampMicro,
	time.StampNano,
}
