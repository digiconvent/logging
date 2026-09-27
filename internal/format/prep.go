package format

import (
	"fmt"
	"strings"
)

func prep(input any) string {
	stringInput := fmt.Sprint(input)
	segments := strings.Split(stringInput, "\n")
	for i := range segments {
		if segments[i] == "" || i == 0 {
			continue
		}
		segments[i] = strings.Repeat(" ", 21) + segments[i]
	}

	result := segments[0]
	for i := 1; i < len(segments); i++ {
		if segments[i] == "" {
			continue
		}
		result += "\n" + segments[i]
	}

	return result
}
