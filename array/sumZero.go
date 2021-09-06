package array

const finalValue = 0

func SumZero(n int) []int {
	var result = make([]int, n)
	var middle int
	var index int

	if n == 1 {
		result[index] = 0
		return result
	}

	if n%2 == 0 {
		middle = n / 2
	} else {
		middle = (n - 1) / 2

		result[index] = finalValue
		index++
	}

	for i := 1; i <= middle; i++ {
		result[index] = i
		result[index+1] = -i

		index += 2
	}

	return result
}
