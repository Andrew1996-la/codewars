package k7

func solution(str, ending string) bool {
	if len(str) < len(ending) {
		return false
	}

	runeStr := []rune(str)
	endStr := string(runeStr[len(runeStr)-len(ending):])
	if endStr == ending {
		return true
	}

	return false
}
