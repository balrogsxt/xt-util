package xweight

import "fmt"

//暂时都是1:1权重

type WeightedRoundRobin[T comparable] struct {
	items         []T
	currentIdx    int
	currentWeight int
}

func NewWeightedRoundRobin[T comparable](items []T) *WeightedRoundRobin[T] {
	return &WeightedRoundRobin[T]{
		items:         items,
		currentIdx:    0,
		currentWeight: 1,
	}
}
func (w *WeightedRoundRobin[T]) List() []T {
	return w.items
}
func (w *WeightedRoundRobin[T]) Add(v T) *WeightedRoundRobin[T] {
	w.items = append(w.items, v)
	return w
}
func (w *WeightedRoundRobin[T]) Remove(v T) *WeightedRoundRobin[T] {
	for i, item := range w.items {
		if item == v {
			w.items = append(w.items[:i], w.items[i+1:]...)
			break
		}
	}
	return w
}

func (w *WeightedRoundRobin[T]) Next() (T, error) {
	if len(w.items) == 0 {
		var def T
		return def, fmt.Errorf("no item")
	}
	item := w.items[w.currentIdx]
	w.currentWeight--

	if w.currentWeight == 0 {
		w.currentIdx = (w.currentIdx + 1) % len(w.items)
		w.currentWeight = 1
	}

	return item, nil
}
