package array

func RepeatedNTimes(nums []int) int {
	elements := make(map[int]int)
	var result int

	for _, number := range nums {
		elements[number] += 1

		if elements[number] == len(nums)/2 {
			result = number
		}
	}

	return result
}
