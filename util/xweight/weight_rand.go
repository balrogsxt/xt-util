package xweight

import (
	"fmt"
	"math/rand"
	"time"
)

// WeightRandomRoundRobin 加权随机
type WeightRandomRoundRobin[T comparable] struct {
	items       []T   //选项列表
	weights     []int //权重
	cumWeights  []int
	totalWeight int //权重总和
}

func NewWeightRandomRoundRobin[T comparable](items []T) *WeightRandomRoundRobin[T] {
	//填充基础权重
	weights := make([]int, len(items))
	for i := range weights {
		weights[i] = 1
	}
	//初始化权重
	w := &WeightRandomRoundRobin[T]{
		items:   items,
		weights: weights,
	}
	w.cumWeights = w.cumulativeWeights(w.weights)
	w.totalWeight = w.cumWeights[len(w.cumWeights)-1]
	return w
}

func (w *WeightRandomRoundRobin[T]) cumulativeWeights(weights []int) []int {
	cumWeights := make([]int, len(weights))
	total := 0
	for i, weight := range weights {
		total += weight
		cumWeights[i] = total
	}
	return cumWeights
}
func (w *WeightRandomRoundRobin[T]) List() []T {
	return w.items
}

// Add 增加选项
func (w *WeightRandomRoundRobin[T]) Add(v T) *WeightRandomRoundRobin[T] {
	w.items = append(w.items, v)
	w.weights = append(w.weights, 1)
	w.cumWeights = w.cumulativeWeights(w.weights)
	w.totalWeight = w.cumWeights[len(w.cumWeights)-1]
	return w
}

// Remove 删除选项
func (w *WeightRandomRoundRobin[T]) Remove(v T) *WeightRandomRoundRobin[T] {
	for i, item := range w.items {
		if item == v {
			w.items = append(w.items[:i], w.items[i+1:]...)
			w.weights = append(w.weights[:i], w.weights[i+1:]...)
			w.cumWeights = w.cumulativeWeights(w.weights)
			w.totalWeight = w.cumWeights[len(w.cumWeights)-1]
			break
		}
	}
	return w
}

func (w *WeightRandomRoundRobin[T]) Next() (T, error) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(w.totalWeight)
	for i, cumWeight := range w.cumWeights {
		if rnd < cumWeight {
			return w.items[i], nil
		}
	}
	if len(w.items) > 0 {
		return w.items[0], nil //取第一个
	}
	var def T
	return def, fmt.Errorf("no item")
}
