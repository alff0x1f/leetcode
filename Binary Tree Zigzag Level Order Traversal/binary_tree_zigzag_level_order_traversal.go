package Binary_Tree_Zigzag_Level_Order_Traversal

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Level struct {
	nodes       []*TreeNode
	leftToRight bool
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return make([][]int, 0)
	}
	answer := make([][]int, 0)
	level := Level{nodes: []*TreeNode{root}, leftToRight: true}
	for len(level.nodes) > 0 {
		levelVals := make([]int, 0)
		nextLevel := Level{
			nodes:       make([]*TreeNode, 0),
			leftToRight: !level.leftToRight,
		}
		for i, node := range level.nodes {
			if level.leftToRight {
				levelVals = append(levelVals, node.Val)
			} else {
				levelVals = append(levelVals, level.nodes[len(level.nodes)-i-1].Val)
			}
			if node.Left != nil {
				nextLevel.nodes = append(nextLevel.nodes, node.Left)
			}
			if node.Right != nil {
				nextLevel.nodes = append(nextLevel.nodes, node.Right)
			}
		}
		answer = append(answer, levelVals)
		level = nextLevel
	}

	return answer
}
