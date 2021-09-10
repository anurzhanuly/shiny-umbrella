package array

func SortArrayByParity(nums []int) []int {
	oddValues := make([]int, len(nums))
	evenValues := make([]int, len(nums))
	oddIdx := 0
	evenIdx := 0

	for _, value := range nums {
		if value%2 == 0 {
			evenValues[evenIdx] = value
			evenIdx += 1

			continue
		}

		oddValues[oddIdx] = value
		oddIdx += 1
	}

	for i := 0; i < evenIdx; i++ {
		nums[i] = evenValues[i]
	}

	for i := evenIdx; i < len(nums); i++ {
		oddIdx -= 1
		nums[i] = oddValues[oddIdx]
	}

	return nums
}
