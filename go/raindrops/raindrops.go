package raindrops

import (
	"strconv"
	"strings"
)

func Convert(number int) string {
	var output []string

	if number%3 == 0 {
		output = append(output, "Pling")
	}
	if number%5 == 0 {
		output = append(output, "Plang")
	}
	if number%7 == 0 {
		output = append(output, "Plong")
	}

	if len(output) == 0 {
		return strconv.Itoa(number)
	}

	return strings.Join(output[:], "")
}
