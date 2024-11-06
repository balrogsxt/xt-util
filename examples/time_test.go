package examples

import (
	"github.com/balrogsxt/xt-util/os/xtime"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	t.Logf("distance: %s", xtime.New("2024-11-05 14:00:00").Distance(time.Now()))
	t.Logf("format: %s", xtime.New().Format("t"))
}
