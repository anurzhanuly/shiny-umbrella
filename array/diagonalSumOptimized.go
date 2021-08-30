package array

func DiagonalSumOptimized(mat [][]int) int {
	var result int

	for index, row := range mat {
		result += row[index]

		if len(mat)-index-1 != index {
			result += row[index]
		}
	}

	return result
}
