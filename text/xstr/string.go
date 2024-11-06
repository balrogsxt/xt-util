package xstr

func Substr(str string, start, length int) string {
	strs := []rune(str)
	return string(strs[start : start+length])
}
