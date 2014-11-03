package atbash

import (
	"unicode"
)

func Atbash(str string) string {
	var encodedMessage []rune

	for _, r := range str {
		normalizedR := unicode.ToLower(r)
		var count int

		encodedR := encodeRune(normalizedR)

		if encodedR != ' ' {
			count += 1
			if count == 5 {
				encodedMessage = append(encodedMessage, ' ')
				count = 0
			}

			encodedMessage = append(encodedMessage, encodedR)
		}
	}

	return string(encodedMessage)
}

func encodeRune(r rune) rune {
	switch {
	case 'a' <= r && r <= 'z':
		return 25 - r + 97 + 97
	default:
		return r
	}
}
