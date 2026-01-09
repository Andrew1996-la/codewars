package k7

import "strings"

func ShortestWord(sentence string) int {
	wordSlice := strings.Split(sentence, " ")
	shortWordLetterCount := len(wordSlice[0])
	for _, word := range wordSlice {
		if len(word) < shortWordLetterCount {
			shortWordLetterCount = len(word)
		}
	}

	return shortWordLetterCount
}
