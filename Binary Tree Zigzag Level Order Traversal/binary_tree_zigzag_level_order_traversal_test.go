package Binary_Tree_Zigzag_Level_Order_Traversal

import (
	"fmt"
	"testing"
)

func TestZigzagLevelOrder(t *testing.T) {
	var tests = []struct {
		tree   *TreeNode
		answer string
	}{
		{
			tree: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{15, nil, nil},
					Right: &TreeNode{7, nil, nil},
				}},
			answer: "[[3] [20 9] [15 7]]",
		}, {
			tree:   &TreeNode{1, nil, nil},
			answer: "[[1]]",
		}, {
			tree:   nil,
			answer: "[]",
		}, {
			tree: &TreeNode{
				1,
				&TreeNode{2, nil, nil},
				nil,
			},
			answer: "[[1] [2]]",
		},
	}
	for _, tst := range tests {
		answer := zigzagLevelOrder(tst.tree)
		ansStr := fmt.Sprint(answer)
		if ansStr != tst.answer {
			t.Errorf("wrong answer, %s, must be %s", ansStr, tst.answer)
		}
	}

}
