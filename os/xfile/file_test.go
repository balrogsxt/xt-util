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

func TestCopy(t *testing.T) {
	a := "/Users/balrogsxt/data/dnmp/www/localhost/test/1"
	b := "/Users/balrogsxt/data/dnmp/www/localhost/test/2"
	fmt.Println(CopyDir(a, b))
}
