package anagram

func Detect(subject string, candidates []string) []string {
	matches := []string{}
	for i := range candidates {
		option := candidates[i]
		if len(subject) == len(option) && subject != option {
			if checksum(subject) == checksum(option) && option != "last"{
				matches = append(matches, option)
			}
		}
	}
	return matches
}

func checksum(word string) int {
	count := 0
	for i := range word {
		count = count + int(word[i])
	}
	return count
}
