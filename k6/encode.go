package k6

import (
	"strconv"
	"strings"
)

/*
Your message is a string containing space separated words.
You need to encrypt each word in the message using the following rules:
1. The first letter must be converted to its ASCII code.
2. The second letter must be switched with the last letter
3. Keepin' it simple: There are no special characters in the input.
*/
//EncryptThis("Hello") == "72olle"

func EncryptThis(text string) string {
	words := strings.Split(text, " ")
	result := make([]string, 0, len(words))

	for _, word := range words {
		if len(word) == 0 {
			continue
		}
		asciiValue := strconv.Itoa(int(word[0]))

		if len(word) == 1 {
			result = append(result, asciiValue)
			continue
		}

		if len(word) == 2 {
			modifyedWord := asciiValue + string(word[1])
			result = append(result, modifyedWord)
			continue
		}

		secondLetter := word[1]
		lastLetter := word[len(word)-1]
		restLetters := word[2 : len(word)-1]
		modifyedWord := asciiValue + string(lastLetter) + restLetters + string(secondLetter)

		result = append(result, modifyedWord)
	}

	return strings.Join(result, " ")
}
