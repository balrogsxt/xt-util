package xslices

import "testing"

func TestStack(t *testing.T) {

	s := NewStack[string]("1", "2", "3")
	s.Push("4")
	t.Log(s.Pop()) //取出最后一个
	t.Log(s.Pop()) //取出最后一个
	t.Log(s.Pop()) //取出最后一个
	t.Log(s.Pop()) //取出最后一个

}
