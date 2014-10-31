package atbash

import (
	"strings"
)

func Atbash(str string) string {
	var encodedChars []string
	var count int

	chars := strings.Split(str, "")

	for _, char := range chars {
		encodedChar := mapping(char)

		if encodedChar != "" {
			count += 1
			encodedChars = append(encodedChars, encodedChar)
		}

		if count == 5 {
			encodedChars = append(encodedChars, " ")
			count = 0
		}
	}

	return strings.TrimSpace(strings.Join(encodedChars, ""))
}

func mapping(str string) string {
	return map[string]string{
		"a": "z",
		"b": "y",
		"c": "x",
		"d": "w",
		"e": "v",
		"f": "u",
		"g": "t",
		"h": "s",
		"i": "r",
		"j": "q",
		"k": "p",
		"l": "o",
		"m": "n",
		"n": "m",
		"o": "l",
		"p": "k",
		"q": "j",
		"r": "i",
		"s": "h",
		"t": "g",
		"u": "f",
		"v": "e",
		"w": "d",
		"x": "c",
		"y": "b",
		"z": "a",
		"1": "1",
		"2": "2",
		"3": "3",
		"4": "4",
		"5": "5",
		"6": "6",
		"7": "7",
		"8": "8",
		"9": "9",
	}[strings.ToLower(str)]
}
