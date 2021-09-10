package array

import "math"

func ThirdMax(nums []int) int {
	max1 := math.MinInt32
	max2 := math.MinInt32
	max3 := math.MinInt32

	for _, value := range nums {
		if value > max1 {

			if max2 < max1 {
				if max3 < max2 {
					max3 = max2
				}
				max2 = max1
			}
			max1 = value

			continue
		}

	}

	return -1
}
