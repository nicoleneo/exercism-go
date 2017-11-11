package isogram

import (
	"strings"
	"unicode"
)

var letterFreqsMap map[rune]int

func IsIsogram(word string) bool {
	letterFreqsMap := make(map[rune]int)
	word = strings.ToLower(word)
	for _, letterRune := range word {
		if _, ok := letterFreqsMap[letterRune]; ok {
			if unicode.IsLetter(letterRune) {
				return false
			}
		}
		letterFreqsMap[letterRune] = 1
	}
	return true
}
