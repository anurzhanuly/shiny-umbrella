package array

func RepeatedNTimes(nums []int) int {
	elements := make(map[int]int)

	for _, number := range nums {
		elements[number] += 1
		if elements[number] > 1 {
			return number
		}
	}

	return -1
}
