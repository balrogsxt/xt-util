package xfile

import (
	"fmt"
	"testing"
)

func TestScan(t *testing.T) {
	//过滤规则
	//files, err := ScanFiles("/", "", nil)
	//if err != nil {
	//	t.Fatal(err.Error())
	//}
	//for _, f := range files {
	//	fmt.Println(f)
	//}

	files, err := ScanDir("/", "*")
	if err != nil {
		t.Fatalf(err.Error())
	}
	for _, f := range files {
		fmt.Println(f)
	}
}
