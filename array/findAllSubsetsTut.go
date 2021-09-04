package array

import "fmt"

func FindAllSubsets(nums []int) [][]int {
	var result [][]int

	findSubsets(nums, make([]int, len(nums)), 0)

	return result
}

func findSubsets(nums []int, subset []int, index int) {
	if index == len(nums) {
		printer(subset)

		return
	}

	subset[index] = nums[index]
	findSubsets(nums, subset, index+1)

	subset[index] = -1
	findSubsets(nums, subset, index+1)
}

func printer(subset []int) {
	for i := 0; i < len(subset); i++ {
		if subset[i] != -1 {
			fmt.Print(subset[i])
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func bruteForceSubset() {

}
