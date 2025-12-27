package k6

func TwoSum(numbers []int, target int) [2]int {
	mapInt := make(map[int]int)

	res := [2]int{}

	for index, num := range numbers {
		need := target - num
		if prevIndex, ok := mapInt[need]; !ok {
			mapInt[num] = index
		} else {
			res[0] = prevIndex
			res[1] = index
		}

	}
	return res
}
