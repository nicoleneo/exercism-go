package hamming

import (
	"errors"
)

const testVersion = 6

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Both strings must be the same length")
	}
	hammingDist := 0
	for i, _ := range a {
		if a[i] != b[i] {
			hammingDist += 1
		}
	}
	return hammingDist, nil
}
