package array

import "testing"

func TestRepeatedNTimes(t *testing.T) {
	nums := [][]int{
		{1, 2, 3, 3},
		{2, 1, 2, 5, 3, 2},
		{5, 1, 5, 2, 5, 3, 5, 4},
	}

	answer := []int{3, 2, 5}

	for index, array := range nums {
		if RepeatedNTimes(array) != answer[index] {
			t.Errorf("Value %v is not equal to %d", array, answer[index])
		}
	}
}
