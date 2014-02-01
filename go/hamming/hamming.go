package hamming

func Distance(strandA string, strandB string) int {
	distance := 0
	shorter := strandB
	if len(strandA) < len(strandB) {
		shorter = strandA
	}
	for i := range shorter {
		if strandA[i] != strandB[i] {
			distance++
		}
	}
	return distance
}
