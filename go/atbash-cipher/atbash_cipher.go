package atbash

import (
	"unicode"
)

func Atbash(str string) string {
	encodedMessage := encodeMessage(str)
	groupedEncodedMessage := addGroupings(encodedMessage)

	return string(groupedEncodedMessage)
}

func encodeMessage(str string) []rune {
	var encodedMessage []rune

	for _, r := range str {
		normalizedR := unicode.ToLower(r)

		encodedR := encodeRune(normalizedR)

		if encodedR != ' ' {
			encodedMessage = append(encodedMessage, encodedR)
		}
	}

	return encodedMessage
}

func addGroupings(m []rune) []rune {
	var groupedEncodedMessage []rune

	for i, r := range m {
		groupedEncodedMessage = append(groupedEncodedMessage, r)

		if (i+1)%5 == 0 && (i != len(m)-1) {
			groupedEncodedMessage = append(groupedEncodedMessage, ' ')
		}
	}

	return groupedEncodedMessage
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
