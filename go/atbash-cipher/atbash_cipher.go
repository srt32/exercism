package atbash

import (
	"unicode"
)

func Atbash(str string) string {
	var encodedMessage []rune
	var groupedEncodedMessage []rune

	for _, r := range str {
		normalizedR := unicode.ToLower(r)

		encodedR := encodeRune(normalizedR)

		if encodedR != ' ' {
			encodedMessage = append(encodedMessage, encodedR)
		}
	}

	for i, r := range encodedMessage {
		groupedEncodedMessage = append(groupedEncodedMessage, r)

		if (i+1)%5 == 0 && (i != len(encodedMessage)-1) {
			groupedEncodedMessage = append(groupedEncodedMessage, ' ')
		}
	}

	return string(groupedEncodedMessage)
}

func encodeRune(r rune) rune {
	switch {
	case 'a' <= r && r <= 'z':
		return 25 - r + 97 + 97
	case 48 <= r && r <= 57:
		return r
	default:
		return ' '
	}
}
