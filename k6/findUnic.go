package k6

func FindUniq(arr []float32) float32 {
	setNumber := make(map[float32]float32)
	var res float32
	for _, number := range arr {
		setNumber[number]++
	}

	for key, value := range setNumber {
		if value == 1 {
			res = key
		}
	}
	return res
}
