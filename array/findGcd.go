package array

func FindGCD(nums []int) int {
	var result int
	var smallestNumber = nums[0]
	var biggestNumber = nums[0]

	for _, number := range nums {
		if number > biggestNumber {
			biggestNumber = number
		}

		if number < smallestNumber {
			smallestNumber = number
		}
	}

	for i := 1; i <= smallestNumber; i++ {
		if smallestNumber%i == 0 && biggestNumber%i == 0 {
			result = i
		}
	}

	return result
}
