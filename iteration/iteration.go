package iteration

import "strings"

func Repeat(char string, num int) string {

	var repeated strings.Builder
	for range num {
		repeated.WriteString(char)
	}
	return repeated.String()
}
