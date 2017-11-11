package acronym

import (
	"regexp"
	"strings"
)

const testVersion = 3

func Abbreviate(sentence string) string {
	re := regexp.MustCompile(`[-\s]`)
	words := re.Split(sentence, -1)
	var acronym []byte
	for _, word := range words {
		acronym = append(acronym, word[0])
	}
	return strings.ToUpper(string(acronym))
}
