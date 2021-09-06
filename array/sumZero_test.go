package array

import (
	"testing"
)

const targetValue = 0

func TestSumZero(t *testing.T) {

	for i := 1; i <= 1000; i++ {
		result := (SumZero(i))

		if getSum(result) != targetValue {
			t.Errorf("%v's sum is not equal to 0", result)
		}
	}
}

func getSum(nums []int) int {
	result := 0

	for _, numbers := range nums {
		result += numbers
	}

	return result
}
