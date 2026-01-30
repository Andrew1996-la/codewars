package k6

import "strings"

// "hello" => ["Hello", "hEllo", "heLlo", "helLo", "hellO"]
func wave(words string) []string {
	result := make([]string, 0)
	lettersCount := len(words)

	for i := 0; i < lettersCount; i++ {
		if words[i] == ' ' {
			continue
		}

		bigLetter := strings.ToUpper(string(words[i]))
		firstPatch := words[:i]
		lastPatch := words[i+1:]
		word := ""

		if i == 0 {
			word = bigLetter + lastPatch
			result = append(result, word)
			continue
		}

		if i == lettersCount-1 {
			word = firstPatch + bigLetter
			result = append(result, word)
			continue
		}

		word = firstPatch + bigLetter + lastPatch
		result = append(result, word)
	}

	return result
}
