package array

func MinOperations(nums []int) int {
	var result int

	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if diff < 0 {
			result += Abs(diff) + 1
			nums[i] = nums[i] + Abs(diff) + 1
		} else if diff == 0 {
			result += 1
			nums[i] = nums[i] + 1
		}
	}

	return result
}

func Abs(number int) int {
	if number < 0 {
		return -number
	}

	return number
}
