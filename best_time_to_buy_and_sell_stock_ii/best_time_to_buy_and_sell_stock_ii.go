package best_time_to_buy_and_sell_stock_ii

func maxProfit(prices []int) int {
	maxP := 0
	l := 0
	r := 1

	for r < len(prices) {
		profit := prices[r] - prices[l]
		if profit > 0 {
			maxP += profit
		}
		l++
		r++
	}

	return maxP
}
