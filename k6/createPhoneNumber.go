package k6

import (
	"fmt"
	"strconv"
)

func CreatePhoneNumber(numbers [10]uint) string {
	var (
		firstPath  string
		secondPath string
		thirdPath  string
	)

	for index, number := range numbers {
		if index < 3 {
			firstPath += strconv.Itoa(int(number))
		}

		if index >= 3 && index < 6 {
			secondPath += strconv.Itoa(int(number))
		}

		if index >= 6 && index <= 9 {
			thirdPath += strconv.Itoa(int(number))
		}
	}

	return fmt.Sprintf("(%s) %s-%s", firstPath, secondPath, thirdPath)
}
