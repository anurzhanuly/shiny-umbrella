package array

import "testing"

func TestCanBeIncreasing(t *testing.T) {
	nums := [][]int{
		{1, 1, 1},
		{2, 3, 2, 1},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 10, 1, 2},
		{123, 302, 32, 333},
	}

	answer := []bool{false, false, true, false, true}

	for index := range nums {
		if CanBeIncreasing(nums[index]) != answer[index] {
			t.Errorf("array %v is not %v", nums[index], answer[index])
		}
	}
}
