package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	runes := []rune(s)
	ln := len(runes)
	if ln == 0 {
		return "", nil
	}
	if unicode.IsDigit(runes[0]) {
		return "", ErrInvalidString
	}
	var b strings.Builder
	var prev rune
	var prevExists bool

	for i := 0; i < ln; i++ {
		r := runes[i]
		if unicode.IsDigit(r) {
			if !prevExists || unicode.IsDigit(prev) {
				return "", ErrInvalidString
			}
			count := int(r - '0')
			if count > 0 {
				b.WriteString(strings.Repeat(string(prev), count))
			}
			prevExists = false
			continue
		}
		if prevExists {
			b.WriteRune(prev)
		}
		prev = r
		prevExists = true
	}
	if prevExists {
		b.WriteRune(prev)
	}
	return b.String(), nil
}
