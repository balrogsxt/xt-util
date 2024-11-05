package examples

import (
	ytime "github.com/balrogsxt/yuki-tools/time"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	t.Logf("distance: %s", ytime.New("2024-11-05 14:00:00").Distance(time.Now()))
	t.Logf("format: %s", ytime.New().Format("t"))
}
