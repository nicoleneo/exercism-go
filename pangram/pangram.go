package pangram

import (
	"regexp"
	"strings"
)

const testVersion = 2

func IsPangram(sentence string) bool {
	re := regexp.MustCompile(`[-\s]`)
	words := re.Split(sentence, -1)
	alphabet := make(map[rune]int)
	for i := 0; i < 26; i++ {
		letter := rune('a' + i)
		alphabet[letter] = 0
	}
	for _, word := range words {
		word = strings.ToLower(word)
		for _, letter := range word {
			index := rune(letter)
			if _, ok := alphabet[index]; ok {
				alphabet[index] += 1

			}
		}
	}
	for _, count := range alphabet {
		if count == 0 {
			return false
		}
	}
	return true
}
