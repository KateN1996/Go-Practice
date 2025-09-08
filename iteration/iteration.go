package iteration

import "strings"

const repeatCount = 5

func Repeat(char string, num int) string {

	var repeated strings.Builder
	for range num {
		repeated.WriteString(char)
	}
	return repeated.String()
}
