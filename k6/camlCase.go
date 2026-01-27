package k6

import "strings"

func CamelCase(s string) string {
	res := ""
	words := strings.Split(s, " ")
	for _, w := range words {
		res += strings.Title(w)
	}

	return res
}
