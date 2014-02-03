package hamming

func Distance(strandA string, strandB string) (distance int) {
	for i := 0; i < len(strandA) && i < len(strandB); i++ {
		if strandA[i] != strandB[i] {
			distance++
		}
	}
	return
}
