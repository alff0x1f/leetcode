package Best_Time_to_Buy_and_Sell_Stock

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxProfit(prices []int) int {
	// You are given an array prices where prices[i] is the price of a given stock on
	// the ith day.
	//
	// You want to maximize your profit by choosing a single day to buy one stock and
	// choosing a different day in the future to sell that stock.
	//
	// Return the maximum profit you can achieve from this transaction. If you cannot
	// achieve any profit, return 0.
	if len(prices) < 2 {
		return 0
	}

	bestProfit := 0
	sellPrice := prices[len(prices)-1]

	for i := len(prices) - 2; i >= 0; i-- {
		bestProfit = max(bestProfit, sellPrice-prices[i])
		sellPrice = max(sellPrice, prices[i])
	}
	return bestProfit
}
