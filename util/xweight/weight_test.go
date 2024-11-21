package xweight

import (
	"fmt"
	"testing"
)

func TestWeight(t *testing.T) {
	items := []string{"1", "2", "3"}
	wrr := NewWeightedRoundRobin[string](items)
	for i := 0; i < 4; i++ {
		fmt.Println(wrr.Next())
	}
	wrr.Add("4")
	// 输出测试
	for i := 0; i < 4; i++ {
		fmt.Println(wrr.Next())
	}
	//items := []string{"1", "2"}
	//w := NewWeightRoundRobin[string](items)
	//// 输出测试
	//for i := 0; i < 4; i++ {
	//	fmt.Println(w.Next())
	//}
	//w.Add("3")
	//fmt.Println("------")
	//for i := 0; i < 6; i++ {
	//	fmt.Println(w.Next())
	//}
}
