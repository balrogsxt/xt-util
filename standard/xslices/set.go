package xslices

// Difference 计算差集
func Difference[T comparable](a, b []T) []T {
	m := make(map[T]bool)
	for _, item := range b {
		m[item] = true
	}
	var list []T
	for _, item := range a {
		if !m[item] {
			list = append(list, item)
		}
	}
	return list
}

// Intersection 计算交集
func Intersection[T comparable](a, b []T) []T {
	m := make(map[T]bool)
	for _, item := range a {
		m[item] = true
	}

	var list []T
	for _, item := range b {
		if m[item] {
			list = append(list, item)
			delete(m, item)
		}
	}
	return list
}
