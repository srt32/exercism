package pascal

func Triangle(depth int) [][]int {
	var triangle [][]int

	for i := 1; i <= depth; i++ {
		var newLayer = []int{}

		for element := 1; element <= i; element++ {
			newLayer = append(newLayer, 1)
		}

		triangle = append(triangle, newLayer)
	}

	return triangle
}
