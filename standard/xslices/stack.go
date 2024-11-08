package xslices

import "sync"

type Stack[T any] struct {
	items []T //栈堆列表
	lock  sync.Mutex
}

func NewStack[T any](items ...T) *Stack[T] {
	return &Stack[T]{
		items: items,
		lock:  sync.Mutex{},
	}
}

// Pop 弹出一个元素
func (s *Stack[T]) Pop() T {
	var empty T
	if len(s.items) == 0 {
		return empty
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// Push 入栈
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Get 弹出一个元素,并且立即加入栈尾部
func (s *Stack[T]) Get() T {
	s.lock.Lock()
	defer s.lock.Unlock()
	t := s.Pop()
	s.Push(t)
	return t
}
