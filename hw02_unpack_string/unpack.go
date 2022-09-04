package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var resultBuilder strings.Builder
	var previousRune rune
	var hasPreviousRune bool

	for _, currentRune := range input {
		if unicode.IsDigit(currentRune) {
			if !hasPreviousRune {
				return "", ErrInvalidString
			}
			repetitions, err := strconv.ParseInt(string(currentRune), 10, 0)
			if err != nil {
				return "", ErrInvalidString
			}

			resultBuilder.WriteString(strings.Repeat(string(previousRune), int(repetitions)))
			hasPreviousRune = false
			continue
		} else if hasPreviousRune {
			resultBuilder.WriteRune(previousRune)
		}

		previousRune = currentRune
		hasPreviousRune = true
	}
	if hasPreviousRune {
		resultBuilder.WriteRune(previousRune)
	}
	return resultBuilder.String(), nil
}
