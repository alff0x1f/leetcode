package Invert_Binary_Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	//Given the root of a binary tree, invert the tree, and return its root.
	if root == nil {
		return root
	}
	tmpNode := root.Right
	root.Right = root.Left
	root.Left = tmpNode
	invertTree(root.Right)
	invertTree(root.Left)
	return root
}
