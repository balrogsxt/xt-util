package xslices

// DropDuplicate 删除重复项
func DropDuplicate[T comparable](s []T) []T {
	if len(s) == 0 {
		return s
	}
	m := make(map[T]struct{}, len(s))
	ret := make([]T, 0, len(m))
	for _, v := range s {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			ret = append(ret, v)
		}
	}
	return ret
}
