package array

func FinalPrices(prices []int) []int {
	for i := 0; i < len(prices); i++ {
		var tmp int

		for j := i + 1; j < len(prices); j++ {
			if prices[i] >= prices[j] {
				tmp = prices[i] - prices[j]
				break
			}
		}

		prices[i] = tmp
	}

	return prices
}
