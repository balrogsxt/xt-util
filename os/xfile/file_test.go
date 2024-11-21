package xfile

import (
	"fmt"
	"testing"
)

func TestScan(t *testing.T) {
	//过滤规则
	files, err := ScanFiles("/Users/balrogsxt/Downloads", "*.apk", nil)
	if err != nil {
		t.Fatal(err.Error())
	}
	for _, f := range files {
		fmt.Println(f)
	}
}
