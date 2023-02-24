package IPO

import (
	"math/rand"
)

/*
Suppose LeetCode will start its IPO soon. In order to sell a good price of its shares
to Venture Capital, LeetCode would like to work on some projects to increase its capital
before the IPO. Since it has limited resources, it can only finish at most k distinct
projects before the IPO. Help LeetCode design the best way to maximize its total capital
after finishing at most k distinct projects.

You are given n projects where the ith project has a pure profit profits[i] and a
minimum capital of capital[i] is needed to start it.

Initially, you have w capital. When you finish a project, you will obtain its pure
profit and the profit will be added to your total capital.

Pick a list of at most k distinct projects from given projects to maximize your final
capital, and return the final maximized capital.

The answer is guaranteed to fit in a 32-bit signed integer.
*/
func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	sortByProfits(profits, capital)
	start := 0
	for i := 0; i < k; i++ {
		for j := start; j < len(profits); j++ {
			if profits[j] > 0 {
				start = j
				break
			}
		}
		for j := start; j < len(profits); j++ {
			if w >= capital[j] && profits[j] > 0 {
				w += profits[j]
				profits[j] = -1
				break
			}
			if j == len(profits)-1 {
				// all project too expensive, not enough capital
				return w
			}
		}
	}
	return w
}

func sortByProfits(profits []int, capital []int) {
	// quicksort algo
	if len(profits) < 2 {
		return
	}

	for i := 0; profits[0] == profits[i]; i++ {
		// check, maybe all elements equal
		if i == len(profits)-1 {
			return
		}
	}

	left, right := 0, len(profits)-1

	pivot := rand.Int() % len(profits)
	profits[pivot], profits[right] = profits[right], profits[pivot]
	capital[pivot], capital[right] = capital[right], capital[pivot]

	for i := range profits {
		if profits[i] > profits[right] {
			profits[left], profits[i] = profits[i], profits[left]
			capital[left], capital[i] = capital[i], capital[left]
			left++
		}
	}
	profits[left], profits[right] = profits[right], profits[left]
	capital[left], capital[right] = capital[right], capital[left]

	sortByProfits(profits[:left], capital[:left])
	sortByProfits(profits[left+1:], capital[left+1:])

}
