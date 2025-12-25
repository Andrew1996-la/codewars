package k7

import "strings"

//most trees are blue
//    to
//Most Trees Are Blue

func ToJadenCase(str string) string {
	arr := strings.Fields(str)
	modifiedRes := make([]string, 0, len(arr))

	for _, word := range arr {
		modified := strings.ToUpper(word[:1]) + word[1:]
		modifiedRes = append(modifiedRes, modified)
	}

	return strings.Join(modifiedRes, " ")
}
