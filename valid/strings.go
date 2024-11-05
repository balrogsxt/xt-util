package valid

import "regexp"

func IsRegex(pattern string, value string) bool {
	has, _ := regexp.MatchString(pattern, value)
	return has
}
