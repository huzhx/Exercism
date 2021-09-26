package raindrops

import (
	"fmt"
	"strings"
)

func Convert(number int) string {
	var output []string

	if number%3 != 0 && number%5 != 0 && number%7 != 0 {
		return fmt.Sprintf("%v", number)
	}

	if number%3 == 0 {
		output = append(output, "Pling")
	}
	if number%5 == 0 {
		output = append(output, "Plang")
	}
	if number%7 == 0 {
		output = append(output, "Plong")
	}

	return fmt.Sprintf(strings.Join(output[:], ""))
}
