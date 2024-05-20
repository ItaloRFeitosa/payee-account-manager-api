package domain

import (
	"errors"
	"strings"
	"unicode"
)

var ErrTemperedValue = errors.New("tempered value")

// keepOnlyNumbers returns a string containing only the numeric characters from the input string.
func keepOnlyNumbers(input string) string {
	var builder strings.Builder
	builder.Grow(len(input)) // Preallocate memory based on the input length

	for _, r := range input {
		if unicode.IsDigit(r) {
			builder.WriteRune(r)
		}
	}

	return builder.String()
}
