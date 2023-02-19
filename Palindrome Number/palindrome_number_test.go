package Palindrome_Number

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	//t.Skip()
	var tests = []struct {
		input  int
		output bool
	}{
		{121, true},
		{-121, false},
		{10, false},
	}
	for _, tst := range tests {
		answer := isPalindrome(tst.input)
		if answer != tst.output {
			t.Errorf("Wrong decision (%t) for %d", answer, tst.input)
		}
	}
}

func TestSplitIntToDigits(t *testing.T) {
	var tests = []struct {
		input  int
		output []int
	}{
		{1, []int{1}},
		{100, []int{1, 0, 0}},
		{12345, []int{1, 2, 3, 4, 5}},
	}
	for _, tst := range tests {
		digits := splitIntToDigits(tst.input)
		if len(digits) != len(tst.output) {
			digitsStr := fmt.Sprint(digits)
			t.Errorf("error in %d (%s)", tst.input, digitsStr)
		}
		for i := range tst.output {
			if tst.output[i] != digits[i] {
				digitsStr := fmt.Sprint(digits)
				t.Errorf("error in %d (%s)", tst.input, digitsStr)
			}
		}
	}
}
