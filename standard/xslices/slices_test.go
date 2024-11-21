package xslices

import "testing"

func TestDifference(t *testing.T) {
	//计算差集元素
	t.Log(Difference([]int{1, 2, 3}, []int{2, 3})) //1
	t.Log(Difference([]int{2, 3, 4}, []int{2, 3})) //4
	//计算交集元素
	t.Log(Intersection([]int{1, 2, 4}, []int{1, 2, 3})) //1,2
}
