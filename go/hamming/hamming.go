package hamming

import "fmt"

func Distance(a, b string) (int, error) {
	hammingDist := 0
	if len(a) != len(b) {
		return -1, fmt.Errorf("a and b have different lengths")
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			hammingDist++
		}
	}
	return hammingDist, nil
}
