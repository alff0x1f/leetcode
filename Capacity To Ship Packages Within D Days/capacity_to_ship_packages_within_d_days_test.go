package Capacity_To_Ship_Packages_Within_D_Days

import (
	"fmt"
	"testing"
)

func TestShipDuration(t *testing.T) {
	tests := []struct {
		weights  []int
		capacity int
		days     int
		maxLoad  int
	}{
		{[]int{1, 1, 1}, 1, 3, 1},
		{[]int{5}, 10, 1, 5},
		{[]int{5, 5, 10}, 10, 2, 10},
	}
	for _, tst := range tests {
		days, maxLoad := shipDuration(&tst.weights, tst.capacity)
		if days != tst.days {
			t.Errorf("Wrong durations, must be %d, not %d",
				tst.days, days)
		}
		if maxLoad != tst.maxLoad {
			t.Errorf("Wrong maxLoad, must be %d, not %d",
				tst.maxLoad, maxLoad)
		}
	}
}

func TestShipWithinDays(t *testing.T) {
	tests := []struct {
		weights  []int
		days     int
		capacity int
	}{
		//{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 15},
		//{[]int{3, 2, 2, 4, 1, 4}, 3, 6},
		{[]int{1, 2, 3, 3, 1, 1}, 5, 3},
	}
	for _, tst := range tests {
		answer := shipWithinDays(tst.weights, tst.days)
		if answer != tst.capacity {
			weightsStr := fmt.Sprint(tst.weights)
			t.Errorf("Error, for %s and %d days capacity must be %d, not %d",
				weightsStr, tst.days, tst.capacity, answer)
		}
	}
}
