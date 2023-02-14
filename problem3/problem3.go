package problem3

import (
	"strconv"
	"strings"
)

func RunLengthEncode(word string) string {
	encodedWord := ""
	count := 1
	for i := 0; i < len(word); i++ {
		nextInRange := i+1 <= len(word)-1
		if nextInRange && string(word[i]) == string(word[i+1]) {
			count++
		} else {
			encodedWord = encodedWord + string(word[i]) + strconv.Itoa(count)
			count = 1
		}

	}
	return encodedWord
}

func RunLengthDecode(word string) string {
	var decodedWord string
	for i := 0; i < len(word)-1; i++ {
		count, err := strconv.Atoi(string(word[i+1]))
		if err == nil {
			decodedWord += strings.Repeat(string(word[i]), count)
		}

	}

	return decodedWord
}
