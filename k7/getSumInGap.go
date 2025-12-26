package k7

func GetSumInGap(a, b int) int {
	sum := 0
	var smallest int
	var biggest int

	if a > b {
		biggest = a
		smallest = b
	} else {
		biggest = b
		smallest = a
	}

	for i := smallest; i <= biggest; i++ {
		sum += i
	}
	return sum
}
