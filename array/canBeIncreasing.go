package array

func CanBeIncreasing(nums []int) bool {
	var lastValue = nums[0]
	var valueB int
	var removeNeeded = false

	for i := 1; i < len(nums); i++ {
		if nums[i] > lastValue {
			valueB = lastValue
			lastValue = nums[i]
		} else {
			if !removeNeeded {
				removeNeeded = true

				if nums[i] > valueB {
					lastValue = nums[i]
				}

				continue
			}

			return false
		}
	}

	return true
}
