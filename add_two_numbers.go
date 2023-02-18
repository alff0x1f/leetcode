package leetcode

// ListNode
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func numberToLinkedList(n int) *ListNode {
	startNode := &ListNode{0, nil}
	node := startNode
	for n > 0 {
		node.Val = n % 10
		n = n / 10
		if n > 0 {
			node.Next = &ListNode{0, nil}
			node = node.Next
		}
	}
	return startNode
}

func linkedListToNumber(node *ListNode) int {
	degree := 1
	summ := 0
	for node.Next != nil {
		summ += node.Val * degree
		degree *= 10
		node = node.Next
	}
	summ += node.Val * degree
	return summ
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// You are given two non-empty linked lists representing two non-negative integers.
	// The digits are stored in reverse order, and each of their nodes contains a single
	// digit. Add the two numbers and return the sum as a linked list.
	//
	// You may assume the two numbers do not contain any leading zero, except the number
	// 0 itself.
	node1 := l1
	node2 := l2
	remain := 0
	lAnswer := &ListNode{0, nil}
	node := lAnswer
	for node1 != nil || node2 != nil || remain != 0 {
		val := remain
		remain = 0
		if node1 != nil {
			val += node1.Val
			node1 = node1.Next
		}
		if node2 != nil {
			val += node2.Val
			node2 = node2.Next
		}
		if val >= 10 {
			remain = val / 10
			val = val % 10
		}
		node.Val = val
		if node1 != nil || node2 != nil || remain != 0 {
			node.Next = &ListNode{0, nil}
			node = node.Next
		}
	}
	return lAnswer
}
