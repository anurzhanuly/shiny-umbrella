package array

func DiagonalSum(mat [][]int) int {
	var result int
	primaryIdx := 0
	secondaryIdx := len(mat) - 1
	var ignoredIdx = -1

	if len(mat)%2 == 1 {
		ignoredIdx = secondaryIdx / 2
	}

	for _, row := range mat {
		result += row[primaryIdx]

		if secondaryIdx != ignoredIdx {
			result += row[secondaryIdx]
		}

		primaryIdx += 1
		secondaryIdx -= 1
	}

	return result
}
