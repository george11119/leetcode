package best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	l := 0
	res := 0

	for r := 1; r < len(prices); r++ {
		if prices[l] > prices[r] {
			l = r
		}

		res = max(res, prices[r]-prices[l])
	}

	return res
}
