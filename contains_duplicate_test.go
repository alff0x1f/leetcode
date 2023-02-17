package leetcode

import "testing"

func TestContainsDuplicate(t *testing.T) {
	var tests = []struct {
		nums   []int
		output bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}
	for _, tst := range tests {
		answer := containsDuplicate(tst.nums)
		if answer != tst.output {
			t.Error("Wrong answer for case", tst.nums)
		}
	}
}
