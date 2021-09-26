package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	hammingDist := 0
	if len(a) != len(b) {
		return -1, errors.New("a and b have different lengths")
	}
	for i := range a {
		if a[i] != b[i] {
			hammingDist++
		}
	}
	return hammingDist, nil
}
