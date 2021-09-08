package array

func CanBeIncreasing(nums []int) bool {
	// 1 var storing the last element
	// 2 var store if the index for removal is found
	var lastValue = nums[0]
	var removeNeeded = false
	// compare values
	for i := 1; i < len(nums); i++ {
		if nums[i] > lastValue {
			lastValue = nums[i]
		} else {
			if !removeNeeded {
				removeNeeded = true
				lastValue = nums[i]

				continue
			}

			return false
		}
	}

	return true
}
