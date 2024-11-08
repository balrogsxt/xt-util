package valid

// IsNumber 是否是数字(包括浮点)
func IsNumber(v string) bool {
	return IsRegex("^\\d+(\\.\\d+)?$", v)
}

// IsInt 是否是整数数字
func IsInt(v string) bool {
	return IsRegex("^\\d+$", v)
}
