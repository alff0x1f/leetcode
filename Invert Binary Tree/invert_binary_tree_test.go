package Invert_Binary_Tree

import "testing"

func TestInvertTree1(t *testing.T) {
	ex1 := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{1, nil, nil},
			Right: &TreeNode{3, nil, nil},
		},
		Right: &TreeNode{
			Val:   7,
			Left:  &TreeNode{6, nil, nil},
			Right: &TreeNode{9, nil, nil},
		},
	}
	answer := invertTree(ex1)
	if answer.Val != 4 ||
		answer.Left.Val != 7 ||
		answer.Right.Val != 2 ||
		answer.Left.Left.Val != 9 ||
		answer.Left.Right.Val != 6 ||
		answer.Right.Left.Val != 3 ||
		answer.Right.Right.Val != 1 {
		t.Error("Wrong invert")
	}
}

func TestInvertTree2(t *testing.T) {
	ex2 := &TreeNode{
		Val:   2,
		Left:  &TreeNode{1, nil, nil},
		Right: &TreeNode{3, nil, nil},
	}
	answer := invertTree(ex2)
	if answer.Val != 2 || answer.Left.Val != 3 || answer.Right.Val != 1 {
		t.Error("Wrong invert")
	}
}

func TestInvertTreeEmpty(t *testing.T) {
	answer := invertTree(nil)
	if answer != nil {
		t.Error("Wrong answer for empty tree")
	}
}
