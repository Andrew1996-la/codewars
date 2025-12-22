package k8

import "fmt"

// [1, 2, 3, 4] => 1 * 2 * 3 * 4 = 24
func grow() {
	nums := []int{4, 1, 1, 1, 4}
	result := nums[0]
	for _, num := range nums[1:] {

		result *= num
	}

	fmt.Println(result)
}
