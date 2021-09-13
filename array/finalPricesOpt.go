package array

func FinalPricesOpt(prices []int) []int {
	var stack []int

	for index := range prices {
		for len(stack) != 0 {
			if prices[len(stack)-1] >= prices[index] {
				prices[stack[len(stack)-1]] -= prices[index]

				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}

		stack = append(stack, index)
	}

	return prices
}
