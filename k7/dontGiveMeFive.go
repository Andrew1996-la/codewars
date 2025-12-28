package k7

import (
	"strconv"
	"strings"
)

func DontGiveMeFive(start int, end int) int {
	res := make([]int, 0)

	for i := start; i <= end; i++ {
		str := strconv.Itoa(i)
		if strings.Contains(str, "5") {
			continue
		}

		res = append(res, i)
	}

	return len(res)
}
