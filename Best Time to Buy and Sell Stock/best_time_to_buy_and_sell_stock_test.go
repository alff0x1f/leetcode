package Best_Time_to_Buy_and_Sell_Stock

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices    []int
		maxProfit int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{1}, 0},
		{[]int{2, 3}, 1},
	}
	for _, test := range tests {
		profit := maxProfit(test.prices)
		if profit != test.maxProfit {
			pricesStr := fmt.Sprint(test.prices)
			t.Errorf("Max profit for %s must be %d, not %d",
				pricesStr, test.maxProfit, profit)
		}
	}
}
