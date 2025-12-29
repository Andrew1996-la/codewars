package k7

import "strings"

/*
i: Increment the value
d: Decrement the value
s: Square the value
o: Output the value to a result array
*/
func Parse(data string) []int {
	res := make([]int, 0)
	initValue := 0
	comands := strings.Split(data, "")
	for _, comand := range comands {
		switch comand {
		case "i":
			initValue++
		case "d":
			initValue--
		case "s":
			initValue *= initValue
		case "o":
			res = append(res, initValue)
		}
	}
	return res
}
