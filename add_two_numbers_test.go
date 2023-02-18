package leetcode

import (
	"testing"
)

func TestNumberToLinkedList(t *testing.T) {
	var tests = []int{5, 80, 123, 12345, 90, 10}
	for _, n := range tests {
		node := numberToLinkedList(n)
		answer := linkedListToNumber(node)
		if answer != n {
			t.Errorf("Wrong sum for %d; value from func - %d", n, answer)
		}
	}
}

func TestAddTwoNumbers(t *testing.T) {
	var tests = []struct {
		numA int
		numB int
	}{
		{342, 465},
		{0, 0},
		{9999, 9999999},
		{9999, 1},
	}
	for _, tst := range tests {
		listA := numberToLinkedList(tst.numA)
		listB := numberToLinkedList(tst.numB)
		answer := addTwoNumbers(listA, listB)
		answerInt := linkedListToNumber(answer)
		if tst.numA+tst.numB != answerInt {
			t.Errorf("Wrong answer, %d + %d != %d", tst.numA, tst.numB, answerInt)
		}
	}
}
