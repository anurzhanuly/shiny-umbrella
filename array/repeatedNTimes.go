package array

func RepeatedNTimes(nums []int) int {
	elements := make(map[int]int)

	for _, number := range nums {
		elements[number] += 1
	}

	for key, value := range elements {
		if value == len(nums)/2 {
			return key
		}
	}

	return -1
}
