package array

func SubsetXORSum(nums []int) int {
	return find(nums, 0, 0)
}

func find(nums []int, index int, xorSum int) int {
	if index == len(nums) {
		return xorSum
	}

	return find(nums, index+1, xorSum^nums[index]) + find(nums, index+1, xorSum)
}
