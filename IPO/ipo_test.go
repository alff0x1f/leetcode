package IPO

import (
	"fmt"
	"testing"
)

func TestFindMaximizedCapital(t *testing.T) {
	tests := []struct {
		k            int
		w            int
		profits      []int
		capital      []int
		finalCapital int
	}{
		{2, 0, []int{1, 2, 3}, []int{0, 1, 1}, 4},
		{3, 0, []int{1, 2, 3}, []int{0, 1, 2}, 6},
		{3, 0, []int{1, 1, 1}, []int{0, 1, 2}, 3},
	}
	for _, tst := range tests {
		answer := findMaximizedCapital(tst.k, tst.w, tst.profits, tst.capital)
		if answer != tst.finalCapital {
			profitStr := fmt.Sprint(tst.profits, tst.capital)
			t.Errorf("Wrong answer (%d) for %s, must be %d",
				answer, profitStr, tst.finalCapital)
		}
	}
}
