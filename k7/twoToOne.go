package k7

import (
	"fmt"
	"sort"
)

func TwoToOne() {
	s1 := "xyaabbbccccdefww"
	s2 := "xxxxyyyyabklmopq"

	set := make(map[rune]bool)

	for _, char := range s1 + s2 {
		set[char] = true
	}

	slice := make([]rune, 0, len(set))

	for char := range set {
		slice = append(slice, char)
	}

	sort.Slice(slice, func(i, b int) bool {
		return slice[i] < slice[b]
	})

	fmt.Println(string(slice))
}
