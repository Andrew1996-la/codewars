package k6

/*
Given an array of integers, find the one that appears an odd number of times.
There will always be only one integer that appears an odd number of times.
*/
func FindOdd(seq []int) int {
	numbersMap := make(map[int]int)

	for _, num := range seq {
		numbersMap[num]++
	}

	for number, howOften := range numbersMap {
		if howOften % 2 == 1 {
			return  number
		}
	}

    return 0
}
