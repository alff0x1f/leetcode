package Minimize_Deviation_in_Array

import (
	"fmt"
	"testing"
)

func TestMinimumDeviation(t *testing.T) {
	tests := []struct {
		nums      []int
		deviation int
	}{
		{nums: []int{1, 2, 3, 4}, deviation: 1},
		{nums: []int{4, 1, 5, 20, 3}, deviation: 3},
		{nums: []int{2, 10, 8}, deviation: 3},
		{nums: []int{1, 3, 5}, deviation: 3},
		{nums: []int{8, 10, 21, 23}, deviation: 15},
		{nums: []int{3, 5}, deviation: 1},
		{nums: []int{10, 4, 3}, deviation: 2},
	}
	for _, tst := range tests {
		origNums := make([]int, len(tst.nums))
		for i := 0; i < len(tst.nums); i++ {
			origNums[i] = tst.nums[i]
		}
		deviation := minimumDeviation(tst.nums)
		numsStr := fmt.Sprint(origNums)
		if deviation != tst.deviation {
			t.Errorf("wrong answer for %s, minimum deviation %d, not %d",
				numsStr, tst.deviation, deviation)
		}
	}
}

func TestSortLast(t *testing.T) {
	tests := []struct {
		arr    []int
		newPos int
	}{
		{[]int{1, 2, 3, 4}, 3},
		{[]int{1, 5, 3}, 1},
		{[]int{2, 10, 8}, 1},
	}
	for _, tst := range tests {
		lastElement := tst.arr[len(tst.arr)-1]
		sortLast(tst.arr)
		if tst.arr[tst.newPos] != lastElement {
			arrStr := fmt.Sprint(tst.arr)
			t.Errorf("Wrong sort: %s", arrStr)
		}
	}
}
