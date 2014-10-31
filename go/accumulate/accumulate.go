package accumulate

func Accumulate(collection []string, converter func(string) string) []string {
	newCollection := []string{}

	for _, str := range collection {
		newCollection = append(newCollection, converter(str))
	}

	return newCollection
}
