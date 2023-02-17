package leetcode

import "testing"

func TestTwoSumBasic(t *testing.T) {
	var tests = []struct {
		nums   []int
		target int
	}{
		{[]int{2, 7, 11, 15}, 9},
		{[]int{3, 2, 4}, 6},
		{[]int{3, 3}, 6},
	}
	for _, tt := range tests {
		ans := twoSum(tt.nums, tt.target)
		if len(ans) != 2 {
			t.Error("answer must be array with two values")
		}
		if tt.nums[ans[0]]+tt.nums[ans[1]] != tt.target {
			t.Errorf("wrong answer, sum (%d+%d) not equal target(%d)",
				tt.nums[ans[0]], tt.nums[ans[1]], tt.target,
			)
		}
	}
}
