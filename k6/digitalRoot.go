package k6

import (
	"strconv"
	"strings"
)

/*
При заданном n возьмите сумму цифр из n. Если это значение содержит более одной цифры,
продолжайте уменьшать таким образом, пока не получится однозначное число. Входными данными
будет неотрицательное целое число.
*/

// 942  -->  9 + 4 + 2 = 15  -->  1 + 5 = 6
func DigitalRoot(n int) int {
	if len(strconv.Itoa(n)) == 1 {
		return n
	}

	numbers := strings.Split(strconv.Itoa(n), "")

	sum := 0
	for _, number := range numbers {
		digit, _ := strconv.Atoi(number)
		sum += digit
	}

	return DigitalRoot(sum)
}
