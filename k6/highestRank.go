package k6

func HighestRank(nums []int) int {
	numMap := make(map[int]int)
	for _, value := range nums {
		numMap[value]++
	}

	maxFrequency := 0
	maxFrequencyValue := -1

	for value, frequency := range numMap {
		if frequency > maxFrequency || (frequency == maxFrequency && value > maxFrequencyValue) {
			maxFrequency = frequency
			maxFrequencyValue = value
		}
	}

	return maxFrequencyValue
}
