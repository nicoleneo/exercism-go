package bob

import (
	"regexp"
	"strings"
)

const testVersion = 3

func Hey(phrase string) string {
	containsAlpha := regexp.MustCompile("[a-zA-Z]+").FindString
	if containsAlpha(phrase) != "" && strings.ToUpper(phrase) == phrase {
		return "Whoa, chill out!"
	}
	if strings.HasSuffix(strings.TrimSpace(phrase), "?") {
		return "Sure."
	}

	if phrase == "" || strings.TrimSpace(phrase) == "" {
		return "Fine. Be that way!"
	}
	return "Whatever."
}
