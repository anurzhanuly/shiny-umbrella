package array

func OddCellsOptimized(m int, n int, indices [][]int) int {
	rowIndices := make(map[int]int)
	columnIndices := make(map[int]int)
	var result int

	for _, value := range indices {
		rowIndices[value[0]] += 1
		columnIndices[value[1]] += 1
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (rowIndices[i]+columnIndices[j])%2 == 1 {
				result += 1
			}
		}
	}

	return result
}
