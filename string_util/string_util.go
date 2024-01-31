package stringutil

import "strings"

func IsEmptyString(input string) bool {
	trimStr := strings.TrimSpace(input)
	return len(trimStr) == 0
}
