package etl

import "strings"

func Transform(input map[int][]string) (out map[string]int) {
	out = make(map[string]int)
	for letter, counts := range input {
		for i := range counts {
			lower_letter := strings.ToLower(counts[i])
			out[lower_letter] = letter
		}
	}
	return out
}
