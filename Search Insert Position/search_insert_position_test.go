package Search_Insert_Position

import "testing"

func TestName(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		output int
	}{
		{
			nums:   []int{1, 3, 5, 6},
			target: 5,
			output: 2,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 2,
			output: 1,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 7,
			output: 4,
		},
		{
			nums:   []int{1, 3},
			target: 2,
			output: 1,
		},
	}
	for _, tst := range tests {
		position := searchInsert(tst.nums, tst.target)
		if position != tst.output {
			t.Errorf("Wrong answer for target %d, position must be %d, not %d",
				tst.target, tst.output, position)
		}
	}
}
