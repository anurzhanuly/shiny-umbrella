package array

func OddCells(m int, n int, indices [][]int) int {
	matrix := make([][]int, m)
	allocateMatrix(n, &matrix)

	var rowIndices = getIndices(indices, 0)
	var columnIndices = getIndices(indices, 1)

	incrementMatrixColumn(&matrix, columnIndices, m)
	incrementMatrixRow(&matrix, rowIndices, n)

	return countOddValues(&matrix, m, n)
}

func allocateMatrix(n int, matrix *[][]int) {

	for index := range *matrix {
		(*matrix)[index] = make([]int, n)
	}
}

func getIndices(indices [][]int, dimension int) map[int]int {
	result := make(map[int]int)

	for _, indicesSet := range indices {
		value := indicesSet[dimension]

		_, ok := result[value]
		if ok {
			result[value] += 1
		} else {
			result[value] = 1
		}

		continue

	}

	return result
}

func incrementMatrixRow(matrix *[][]int, indices map[int]int, colLen int) {
	for key := range indices {
		for j := 0; j < colLen; j++ {
			(*matrix)[key][j] += indices[key]
		}
	}
}

func incrementMatrixColumn(matrix *[][]int, indices map[int]int, rowLen int) {
	for key := range indices {
		for j := 0; j < rowLen; j++ {
			(*matrix)[j][key] += indices[key]
		}
	}
}

func countOddValues(matrix *[][]int, rowLen, colLen int) int {
	result := 0

	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			if (*matrix)[i][j]%2 == 1 {
				result += 1
			}
		}
	}

	return result
}
