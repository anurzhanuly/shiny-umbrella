package array

func OddCells(m int, n int, indices [][]int) int {
	matrix := make([][]int, m)
	allocateMatrix(n, &matrix)

	var rowIncrements = getRowIndices(indices)
	var columnIncrements = getColumnIndices(indices)

}

func allocateMatrix(n int, matrix *[][]int) {

	for index := range *matrix {
		(*matrix)[index] = make([]int, n)
	}
}

func getRowIndices(indices [][]int) map[int]int {
	rowIndices := make(map[int]int)

	for _, indicesSet := range indices {
		for _, value := range indicesSet {
			_, ok := rowIndices[value]
			if ok {
				rowIndices[value] += 1
			} else {
				rowIndices[value] = 1
			}

			continue
		}
	}

	return rowIndices
}
