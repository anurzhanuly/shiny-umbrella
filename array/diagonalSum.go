package array

func DiagonalSum(mat [][]int) int {
	var result int
	//2: I can make 2 pointers, 1 starts from 0 to len(n-1) and 2nd one is opposite
	primaryIdx := 0
	secondaryIdx := len(mat) - 1
	var ignoredIdx = -1

	//1: find the length of the row to understand whether center is ignored
	if secondaryIdx%2 == 1 {
		//1.5: if the row is odd make block for center element
		ignoredIdx = secondaryIdx / 2
	}
	//3: Iterate over the rows and sum all of it
	for _, row := range mat {
		result += row[primaryIdx]

		if secondaryIdx != ignoredIdx {
			result += row[secondaryIdx]
		}

		primaryIdx += 1
		secondaryIdx += 1
	}

	return result
}
