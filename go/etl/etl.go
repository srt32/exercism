package etl

import "strings"

func Transform(input map[int][]string) (output map[string]int) {
	output = make(map[string]int)
	for letter, scores := range input {
		for _, score := range scores {
			lower_letter := strings.ToLower(score)
			output[lower_letter] = letter
		}
	}
	return output
}
