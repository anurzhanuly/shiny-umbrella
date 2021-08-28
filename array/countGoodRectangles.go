package array

func CountGoodRectangles(rectangles [][]int) int {
	squareValues := make(map[int]int)
	maxLen := 0

	for _, value := range rectangles {
		tmp := getLowestValue(value)
		if tmp > maxLen {
			maxLen = tmp
		}

		squareValues[tmp] += 1
	}

	return squareValues[maxLen]
}

func getLowestValue(nums []int) int {
	if nums[0] < nums[1] {
		return nums[0]
	}

	return nums[1]
}
